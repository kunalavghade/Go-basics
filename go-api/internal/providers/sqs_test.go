package providers_test

import (
	"testing"

	"github.com/kunalavghade/Go-basics/go-api/internal/config"
	"github.com/kunalavghade/Go-basics/go-api/internal/providers"
)

func TestSQSEventPublisher(t *testing.T) {
	cfg := &config.Config{
		AWS: config.AWSConfig{
			Region:          "us-east-1",
			AccessKeyID:     "test",
			SecretAccessKey: "test",
			S3Endpoint:      "http://localhost:4566",
			EventQueue:      "task-queue",
		},
		Event: config.EventConfig{
			Provider: "sqs",
			Queue:    "task-queue",
		},
	}

	publisher, err := providers.NewEventPublisher(cfg)
	if err != nil {
		t.Fatalf("failed to create EventPublisher via factory: %v", err)
	}
	defer publisher.Close()

	payload := map[string]interface{}{
		"order_id": 123,
		"status":   "created",
		"total":    99.99,
	}
	metadata := map[string]string{
		"user_id":  "42",
		"trace_id": "test-trace-123",
	}

	err = publisher.Publish("order.created", payload, metadata)
	if err != nil {
		t.Fatalf("failed to publish event: %v", err)
	}
}

func TestEventPublisherProviderSwitching(t *testing.T) {
	// Test switching to noop provider
	cfgNoop := &config.Config{
		Event: config.EventConfig{
			Provider: "noop",
		},
	}
	noopPub, err := providers.NewEventPublisher(cfgNoop)
	if err != nil {
		t.Fatalf("failed to create noop publisher: %v", err)
	}
	defer noopPub.Close()

	if err := noopPub.Publish("order.created", map[string]string{"foo": "bar"}, nil); err != nil {
		t.Fatalf("noop publish failed: %v", err)
	}
}
