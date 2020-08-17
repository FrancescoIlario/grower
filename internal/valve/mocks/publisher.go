package mocks

import (
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
)

// mockPublisher mock that implements the watermill message.Publisher interface
type mockPublisher struct {
	publish func(topic string, messages ...*message.Message) error
	close   func() error
}

// DefaultPublisher MockPublisher constructor
func DefaultPublisher() message.Publisher {
	return &mockPublisher{}
}

// CustomPublisher programmable MockPublisher constructor
func CustomPublisher(publish func(topic string, messages ...*message.Message) error) message.Publisher {
	return &mockPublisher{
		publish: publish,
		close: func() error {
			return nil
		},
	}
}

// CustomClosePublisher programmable MockPublisher constructor with close function
func CustomClosePublisher(publish func(topic string, messages ...*message.Message) error, close func() error) message.Publisher {
	return &mockPublisher{
		publish: publish,
		close:   close,
	}
}

// Publish mocks the publish of the message
func (p *mockPublisher) Publish(topic string, messages ...*message.Message) error {
	if p.publish != nil {
		return p.publish(topic, messages...)
	}

	for _, msg := range messages {
		log.Printf("message published: %v", msg.UUID)
	}
	return nil
}

// Close disposes the MockPublisher
func (p *mockPublisher) Close() error {
	if p.close != nil {
		return p.close()
	}
	return nil
}
