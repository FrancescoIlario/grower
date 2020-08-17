package mocks

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

// mockSubscriber mock that implements the watermill message.Subscriber interface
type mockSubscriber struct {
	subscribe func(ctx context.Context, topic string) (<-chan *message.Message, error)
	close     func() error
}

// DefaultSubscriber MockSubscriber constructor
func DefaultSubscriber() message.Subscriber {
	return &mockSubscriber{}
}

// CustomSubscriber programmable MockSubscriber constructor
func CustomSubscriber(subscribe func(ctx context.Context, topic string) (<-chan *message.Message, error)) message.Subscriber {
	return &mockSubscriber{
		subscribe: subscribe,
		close: func() error {
			return nil
		},
	}
}

// CustomCloseSubscriber programmable MockSubscriber constructor with close function
func CustomCloseSubscriber(subscribe func(ctx context.Context, topic string) (<-chan *message.Message, error), close func() error) message.Subscriber {
	return &mockSubscriber{
		subscribe: subscribe,
		close:     close,
	}
}

// Subscribe mocks the subscribe of the message
func (p *mockSubscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	if p.subscribe != nil {
		return p.subscribe(ctx, topic)
	}
	return nil, nil
}

// Close disposes the MockSubscriber
func (p *mockSubscriber) Close() error {
	if p.close != nil {
		return p.close()
	}
	return nil
}
