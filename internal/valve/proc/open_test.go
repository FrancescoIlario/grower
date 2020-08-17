package proc_test

import (
	"context"
	"testing"
	"time"

	"github.com/FrancescoIlario/grower/internal/mocks"
	"github.com/FrancescoIlario/grower/internal/valve/proc"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

func Test_OpenHandler(t *testing.T) {
	publisher := mocks.DefaultPublisher()
	marshaler := cqrs.ProtobufMarshaler{}
	eb, err := cqrs.NewEventBus(publisher, func(eventName string) string { return eventName }, marshaler)
	if err != nil {
		t.Fatalf("error creating the event bus: %v", err)
	}

	cmder := mocks.NewValveCmder(200 * time.Millisecond)
	handler := proc.NewOpenHandler(eb, cmder)

	ctx := context.Background()

	cmd := &proc.OpenCommand{}
	if err := handler.Handle(ctx, cmd); err != nil {
		t.Fatalf("error handling Open Command: %v", err)
	}

	if exp := publisher.PublishCounter(); exp != 1 {
		t.Fatalf("publish counter expected to be 1, obtained %d", exp)
	}

	if exp, obt := 1, cmder.OpenInvokation(); obt != exp {
		t.Errorf("commander open invokation expected to be %d, obtained %d", exp, obt)
	}
	if exp, obt := 0, cmder.CloseInvokation(); obt != exp {
		t.Errorf("commander close invokation expected to be %d, obtained %d", exp, obt)
	}
}
