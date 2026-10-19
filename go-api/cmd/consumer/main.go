package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/joho/godotenv"
	"github.com/kunalavghade/Go-basics/go-api/internal/providers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	SQSEndpoint string
	QueueName   string
	Region      string
	SMTPHost    string
	SMTPPort    string
	FromEmail   string
}

func loadConfig() *Config {
	_ = godotenv.Load()

	endpoint := getEnv("AWS_S3_ENDPOINT", "http://localhost:4566")
	queue := getEnv("AWS_EVENT_QUEUE", "task-queue")
	region := getEnv("AWS_REGION", "us-east-1")
	smtpHost := getEnv("SMTP_HOST", "localhost")
	smtpPort := getEnv("SMTP_PORT", "1025")
	fromEmail := getEnv("FROM_EMAIL", "noreply@ecommerce.local")

	return &Config{
		SQSEndpoint: endpoint,
		QueueName:   queue,
		Region:      region,
		SMTPHost:    smtpHost,
		SMTPPort:    smtpPort,
		FromEmail:   fromEmail,
	}
}

func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

type SQSReceiveResponse struct {
	XMLName  xml.Name     `xml:"ReceiveMessageResponse"`
	Messages []SQSMessage `xml:"ReceiveMessageResult>Message"`
}

type SQSMessage struct {
	MessageId     string `xml:"MessageId"`
	ReceiptHandle string `xml:"ReceiptHandle"`
	Body          string `xml:"Body"`
}

type Consumer struct {
	cfg        *Config
	httpClient *http.Client
	queueURL   string
	sqsClient  *sqs.Client
}

func NewConsumer(cfg *Config) (*Consumer, error) {
	httpClient := &http.Client{Timeout: 15 * time.Second}

	c := &Consumer{
		cfg:        cfg,
		httpClient: httpClient,
	}

	if cfg.SQSEndpoint != "" {
		base := strings.TrimRight(cfg.SQSEndpoint, "/")
		c.queueURL = fmt.Sprintf("%s/000000000000/%s", base, cfg.QueueName)
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		awsCfg, err := providers.CreateAWSConfig(ctx, "", cfg.Region)
		if err != nil {
			return nil, fmt.Errorf("failed to init AWS config: %w", err)
		}
		c.sqsClient = sqs.NewFromConfig(awsCfg)
		out, err := c.sqsClient.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
			QueueName: aws.String(cfg.QueueName),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get queue url: %w", err)
		}
		c.queueURL = *out.QueueUrl
	}

	return c, nil
}

func (c *Consumer) Run(ctx context.Context) {
	log.Info().
		Str("queue_url", c.queueURL).
		Str("smtp", fmt.Sprintf("%s:%s", c.cfg.SMTPHost, c.cfg.SMTPPort)).
		Msg("Event consumer service started, waiting for SQS messages...")

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("Consumer shutting down gracefully...")
			return
		default:
			c.pollMessages(ctx)
		}
	}
}

func (c *Consumer) pollMessages(ctx context.Context) {
	if c.cfg.SQSEndpoint != "" {
		c.pollLocalStack(ctx)
	} else {
		c.pollAWS(ctx)
	}
}

func (c *Consumer) pollLocalStack(ctx context.Context) {
	form := url.Values{}
	form.Set("Action", "ReceiveMessage")
	form.Set("WaitTimeSeconds", "5")
	form.Set("MaxNumberOfMessages", "10")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.queueURL, strings.NewReader(form.Encode()))
	if err != nil {
		if ctx.Err() != nil {
			return
		}
		log.Error().Err(err).Msg("Failed to build receive request")
		time.Sleep(2 * time.Second)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return
		}
		log.Warn().Err(err).Msg("Failed to connect to LocalStack SQS, retrying...")
		time.Sleep(3 * time.Second)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read SQS response")
		return
	}

	var xmlResp SQSReceiveResponse
	if err := xml.Unmarshal(bodyBytes, &xmlResp); err != nil {
		return
	}

	for _, msg := range xmlResp.Messages {
		c.processMessage(msg.ReceiptHandle, msg.Body)
	}
}

func (c *Consumer) pollAWS(ctx context.Context) {
	out, err := c.sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(c.queueURL),
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     5,
	})
	if err != nil {
		if ctx.Err() != nil {
			return
		}
		log.Error().Err(err).Msg("Failed to receive messages from AWS SQS")
		time.Sleep(3 * time.Second)
		return
	}

	for _, msg := range out.Messages {
		body := ""
		receipt := ""
		if msg.Body != nil {
			body = *msg.Body
		}
		if msg.ReceiptHandle != nil {
			receipt = *msg.ReceiptHandle
		}
		c.processMessage(receipt, body)
	}
}

func (c *Consumer) processMessage(receiptHandle, body string) {
	var eventMsg providers.EventMessage
	if err := json.Unmarshal([]byte(body), &eventMsg); err != nil {
		log.Warn().Str("body", body).Msg("Received non-JSON message, skipping")
		c.deleteMessage(receiptHandle)
		return
	}

	log.Info().
		Str("event_id", eventMsg.EventID).
		Str("event_type", eventMsg.EventType).
		Msg("Processing domain event")

	if err := c.handleEvent(&eventMsg); err != nil {
		log.Error().Err(err).Str("event_id", eventMsg.EventID).Msg("Failed to send email for event")
		return
	}

	c.deleteMessage(receiptHandle)
}

func (c *Consumer) deleteMessage(receiptHandle string) {
	if c.cfg.SQSEndpoint != "" {
		form := url.Values{}
		form.Set("Action", "DeleteMessage")
		form.Set("ReceiptHandle", receiptHandle)

		resp, err := c.httpClient.PostForm(c.queueURL, form)
		if err == nil {
			_ = resp.Body.Close()
		}
	} else if c.sqsClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, _ = c.sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(c.queueURL),
			ReceiptHandle: aws.String(receiptHandle),
		})
	}
}

func (c *Consumer) handleEvent(event *providers.EventMessage) error {
	switch event.EventType {
	case "order.created":
		return c.handleOrderCreated(event)
	case "user.login":
		return c.handleUserLogin(event)
	default:
		return c.handleGenericEvent(event)
	}
}

func (c *Consumer) handleOrderCreated(event *providers.EventMessage) error {
	toEmail := "customer@example.com"
	if email, ok := event.Metadata["email"]; ok && email != "" {
		toEmail = email
	}

	orderID := event.Metadata["order_id"]
	if pMap, ok := event.Payload.(map[string]interface{}); ok {
		if orderID == "" {
			if id, exists := pMap["id"]; exists {
				orderID = fmt.Sprintf("%v", id)
			} else if id, exists := pMap["order_id"]; exists {
				orderID = fmt.Sprintf("%v", id)
			}
		}
		if toEmail == "customer@example.com" {
			if email, exists := pMap["email"]; exists && email != "" {
				toEmail = fmt.Sprintf("%v", email)
			}
		}
	}

	subject := fmt.Sprintf("Order Confirmation - Order #%s", orderID)

	payloadBytes, _ := json.MarshalIndent(event.Payload, "", "  ")
	body := fmt.Sprintf(
		"Hello!\r\n\r\nThank you for your order!\r\n\r\nOrder ID: %s\r\nEvent Timestamp: %s\r\n\r\nOrder Details:\r\n%s\r\n\r\nWe will notify you when it ships!\r\n",
		orderID,
		event.Timestamp.Format(time.RFC1123),
		string(payloadBytes),
	)

	return c.sendEmail(toEmail, subject, body)
}

func (c *Consumer) handleUserLogin(event *providers.EventMessage) error {
	toEmail := event.Metadata["email"]
	if toEmail == "" {
		toEmail = "user@example.com"
	}

	subject := "Security Alert: New Login to your Account"
	body := fmt.Sprintf(
		"Hello!\r\n\r\nA new login was detected on your account.\r\n\r\nEmail: %s\r\nTimestamp: %s\r\n\r\nIf this was not you, please secure your account immediately.\r\n",
		toEmail,
		event.Timestamp.Format(time.RFC1123),
	)

	return c.sendEmail(toEmail, subject, body)
}

func (c *Consumer) handleGenericEvent(event *providers.EventMessage) error {
	toEmail := "admin@example.com"
	subject := fmt.Sprintf("Domain Event Notification: %s", event.EventType)

	payloadBytes, _ := json.MarshalIndent(event.Payload, "", "  ")
	body := fmt.Sprintf(
		"Event: %s\r\nEvent ID: %s\r\nTimestamp: %s\r\n\r\nPayload:\r\n%s\r\n",
		event.EventType,
		event.EventID,
		event.Timestamp.Format(time.RFC1123),
		string(payloadBytes),
	)

	return c.sendEmail(toEmail, subject, body)
}

func (c *Consumer) sendEmail(to, subject, body string) error {
	addr := fmt.Sprintf("%s:%s", c.cfg.SMTPHost, c.cfg.SMTPPort)
	headers := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n", c.cfg.FromEmail, to, subject)
	msg := []byte(headers + body)

	log.Info().
		Str("to", to).
		Str("subject", subject).
		Str("smtp_addr", addr).
		Msg("Sending email via SMTP to Mailhog")

	err := smtp.SendMail(addr, nil, c.cfg.FromEmail, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email via SMTP (%s): %w", addr, err)
	}

	log.Info().
		Str("to", to).
		Str("subject", subject).
		Msg("Email successfully delivered to Mailhog")

	return nil
}

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	cfg := loadConfig()

	consumer, err := NewConsumer(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize consumer")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Info().Msg("Interrupt received, stopping consumer...")
		cancel()
	}()

	consumer.Run(ctx)
}
