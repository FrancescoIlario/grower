package mocks

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type subscriber struct {
	subscribe        func(ctx context.Context, topic string) (<-chan *message.Message, error)
	close            func() error
	subscribeCounter int
	closeCounter     int
}

// Subscriber mock that overloads the watermill message.Subscriber interface
type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
	Close() error
	SubscribeCounter() int
	CloseCounter() int
}

// DefaultSubscriber MockSubscriber constructor
func DefaultSubscriber() Subscriber {
	return &subscriber{
		subscribeCounter: 0,
		closeCounter:     0,
	}
}

// CustomSubscriber programmable MockSubscriber constructor
func CustomSubscriber(subscribe func(ctx context.Context, topic string) (<-chan *message.Message, error)) Subscriber {
	return &subscriber{
		subscribe: subscribe,
		close: func() error {
			return nil
		},
		subscribeCounter: 0,
		closeCounter:     0,
	}
}

// CustomCloseSubscriber programmable MockSubscriber constructor with close function
func CustomCloseSubscriber(subscribe func(ctx context.Context, topic string) (<-chan *message.Message, error), close func() error) Subscriber {
	return &subscriber{
		subscribe:        subscribe,
		close:            close,
		subscribeCounter: 0,
		closeCounter:     0,
	}
}

// Subscribe mocks the subscribe of the message
func (p *subscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	p.subscribeCounter++
	if p.subscribe != nil {
		return p.subscribe(ctx, topic)
	}
	return nil, nil
}

// Close disposes the MockSubscriber
func (p *subscriber) Close() error {
	p.closeCounter++
	if p.close != nil {
		return p.close()
	}
	return nil
}

func (p *subscriber) CloseCounter() int {
	return p.closeCounter
}

func (p *subscriber) SubscribeCounter() int {
	return p.subscribeCounter
}
