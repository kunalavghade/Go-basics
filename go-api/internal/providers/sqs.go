package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
	appConfig "github.com/kunalavghade/Go-basics/go-api/internal/config"
	"github.com/kunalavghade/Go-basics/go-api/internal/interfaces"
)

type EventMessage struct {
	EventID   string            `json:"event_id"`
	EventType string            `json:"event_type"`
	Timestamp time.Time         `json:"timestamp"`
	Payload   interface{}       `json:"payload"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type SQSEventPublisher struct {
	client     *sqs.Client
	queueURL   string
	endpoint   string
	httpClient *http.Client
}

// NewSQSEventPublisher creates an EventPublisher using AWS SQS, configured via CreateAWSConfig.
func NewSQSEventPublisher(cfg *appConfig.Config) (interfaces.EventPublisher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	queueName := cfg.Event.Queue
	if queueName == "" {
		queueName = cfg.AWS.EventQueue
	}
	if queueName == "" {
		queueName = "task-queue"
	}

	endpoint := cfg.AWS.S3Endpoint
	region := cfg.AWS.Region
	if region == "" {
		region = "us-east-1"
	}

	// Use CreateAWSConfig from providers/aws.go as requested
	awsCfg, err := CreateAWSConfig(ctx, endpoint, region)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS config for SQS: %w", err)
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	// For LocalStack, use direct query API to ensure compatibility across all LocalStack versions
	if endpoint != "" {
		baseEndpoint := strings.TrimRight(endpoint, "/")
		queueURL := fmt.Sprintf("%s/000000000000/%s", baseEndpoint, queueName)

		// Auto-create queue in LocalStack if it does not exist
		form := url.Values{}
		form.Set("Action", "CreateQueue")
		form.Set("QueueName", queueName)
		resp, createErr := httpClient.PostForm(baseEndpoint+"/", form)
		if createErr == nil {
			_ = resp.Body.Close()
		}

		return &SQSEventPublisher{
			queueURL:   queueURL,
			endpoint:   endpoint,
			httpClient: httpClient,
		}, nil
	}

	// Production AWS SQS client from configured aws.Config
	client := sqs.NewFromConfig(awsCfg)
	getOut, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get SQS queue URL for %q: %w", queueName, err)
	}

	return &SQSEventPublisher{
		client:     client,
		queueURL:   *getOut.QueueUrl,
		httpClient: httpClient,
	}, nil
}

func (p *SQSEventPublisher) Publish(eventType string, payload interface{}, metadata map[string]string) error {
	eventMsg := EventMessage{
		EventID:   uuid.New().String(),
		EventType: eventType,
		Timestamp: time.Now().UTC(),
		Payload:   payload,
		Metadata:  metadata,
	}

	bodyBytes, err := json.Marshal(eventMsg)
	if err != nil {
		return fmt.Errorf("failed to marshal event message: %w", err)
	}

	// LocalStack Query path
	if p.endpoint != "" {
		return p.publishLocalStack(eventType, string(bodyBytes), metadata)
	}

	// Standard AWS SDK path
	msgAttrs := make(map[string]types.MessageAttributeValue)
	msgAttrs["EventType"] = types.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(eventType),
	}

	for k, v := range metadata {
		msgAttrs[k] = types.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(v),
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = p.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:          aws.String(p.queueURL),
		MessageBody:       aws.String(string(bodyBytes)),
		MessageAttributes: msgAttrs,
	})
	if err != nil {
		return fmt.Errorf("failed to send message to SQS: %w", err)
	}

	return nil
}

func (p *SQSEventPublisher) publishLocalStack(eventType, messageBody string, metadata map[string]string) error {
	data := url.Values{}
	data.Set("Action", "SendMessage")
	data.Set("MessageBody", messageBody)

	attrIdx := 1
	data.Set(fmt.Sprintf("MessageAttribute.%d.Name", attrIdx), "EventType")
	data.Set(fmt.Sprintf("MessageAttribute.%d.Value.DataType", attrIdx), "String")
	data.Set(fmt.Sprintf("MessageAttribute.%d.Value.StringValue", attrIdx), eventType)

	for k, v := range metadata {
		attrIdx++
		data.Set(fmt.Sprintf("MessageAttribute.%d.Name", attrIdx), k)
		data.Set(fmt.Sprintf("MessageAttribute.%d.Value.DataType", attrIdx), "String")
		data.Set(fmt.Sprintf("MessageAttribute.%d.Value.StringValue", attrIdx), v)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.queueURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to build SQS request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send SQS request to LocalStack: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("SQS LocalStack returned HTTP status %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}

func (p *SQSEventPublisher) Close() error {
	return nil
}

// NoopEventPublisher provides a no-op fallback implementation of interfaces.EventPublisher
type NoopEventPublisher struct{}

func NewNoopEventPublisher() interfaces.EventPublisher {
	return &NoopEventPublisher{}
}

func (n *NoopEventPublisher) Publish(eventType string, payload interface{}, metadata map[string]string) error {
	return nil
}

func (n *NoopEventPublisher) Close() error {
	return nil
}

// NewEventPublisher creates an EventPublisher based on cfg.Event.Provider.
// This allows easily changing the provider (e.g., SQS, Kafka, RabbitMQ, Noop) anytime.
func NewEventPublisher(cfg *appConfig.Config) (interfaces.EventPublisher, error) {
	provider := strings.ToLower(cfg.Event.Provider)
	if provider == "" {
		provider = "sqs"
	}

	switch provider {
	case "sqs", "aws":
		return NewSQSEventPublisher(cfg)
	case "noop", "none":
		return NewNoopEventPublisher(), nil
	default:
		return nil, fmt.Errorf("unsupported event publisher provider: %q", cfg.Event.Provider)
	}
}
