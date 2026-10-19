package interfaces

// EventPublisher defines the contract for publishing domain events.
type EventPublisher interface {
	Publish(eventType string, payload interface{}, metadata map[string]string) error
	Close() error
}
