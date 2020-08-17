package mocks

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

// publisher mock that implements the watermill message.Publisher interface
type publisher struct {
	publish           func(topic string, messages ...*message.Message) error
	close             func() error
	publishedMessages []*message.Message
	closeCounter      int
}

// Publisher ...
type Publisher interface {
	message.Publisher

	CloseCounter() int
	PublishCounter() int
	PublishedMessages() []*message.Message
}

// DefaultPublisher MockPublisher constructor
func DefaultPublisher() Publisher {
	return &publisher{
		publishedMessages: []*message.Message{},
		closeCounter:      0,
	}
}

// CustomPublisher programmable MockPublisher constructor
func CustomPublisher(publish func(topic string, messages ...*message.Message) error) Publisher {
	return &publisher{
		publish: publish,
		close: func() error {
			return nil
		},
		publishedMessages: []*message.Message{},
		closeCounter:      0,
	}
}

// CustomClosePublisher programmable MockPublisher constructor with close function
func CustomClosePublisher(publish func(topic string, messages ...*message.Message) error, close func() error) Publisher {
	return &publisher{
		publish:           publish,
		close:             close,
		publishedMessages: []*message.Message{},
		closeCounter:      0,
	}
}

// Publish mocks the publish of the message
func (p *publisher) Publish(topic string, messages ...*message.Message) error {
	p.publishedMessages = append(p.publishedMessages, messages...)
	if p.publish != nil {
		return p.publish(topic, messages...)
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
	return len(p.publishedMessages)
}

func (p *publisher) PublishedMessages() []*message.Message {
	return p.publishedMessages
}
