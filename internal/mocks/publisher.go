package mocks

import (
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
)

// publisher mock that implements the watermill message.Publisher interface
type publisher struct {
	publish        func(topic string, messages ...*message.Message) error
	close          func() error
	publishCounter int
	closeCounter   int
}

type Publisher interface {
	Publish(topic string, messages ...*message.Message) error
	Close() error
	CloseCounter() int
	PublishCounter() int
}

// DefaultPublisher MockPublisher constructor
func DefaultPublisher() Publisher {
	return &publisher{
		publishCounter: 0,
		closeCounter:   0,
	}
}

// CustomPublisher programmable MockPublisher constructor
func CustomPublisher(publish func(topic string, messages ...*message.Message) error) Publisher {
	return &publisher{
		publish: publish,
		close: func() error {
			return nil
		},
		publishCounter: 0,
		closeCounter:   0,
	}
}

// CustomClosePublisher programmable MockPublisher constructor with close function
func CustomClosePublisher(publish func(topic string, messages ...*message.Message) error, close func() error) Publisher {
	return &publisher{
		publish:        publish,
		close:          close,
		publishCounter: 0,
		closeCounter:   0,
	}
}

// Publish mocks the publish of the message
func (p *publisher) Publish(topic string, messages ...*message.Message) error {
	p.publishCounter++
	if p.publish != nil {
		return p.publish(topic, messages...)
	}

	for _, msg := range messages {
		log.Printf("message published: %v", msg.UUID)
	}
	return nil
}

// Close disposes the MockPublisher
func (p *publisher) Close() error {
	p.closeCounter++
	if p.close != nil {
		return p.close()
	}
	return nil
}

func (p *publisher) CloseCounter() int {
	return p.closeCounter
}

func (p *publisher) PublishCounter() int {
	return p.publishCounter
}
