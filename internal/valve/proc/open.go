package proc

import (
	"context"

	"github.com/FrancescoIlario/grower/pkg/valvepb/es"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
)

// OpenCommand ...
type OpenCommand struct{}

// OpenedEvent ...
type OpenedEvent struct{}

// OpenHandler is a command handler, which handles OpenCommand and emits ValveOpened.
//
// In CQRS, one command must be handled by only one handler.
// When another handler with this command is added to command processor, error will be returned.
type OpenHandler struct {
	eventBus *cqrs.EventBus
	Cmder    Commander
}

// HandlerName ...
func (b OpenHandler) HandlerName() string {
	return "OpenHandler"
}

// NewCommand returns type of command which this handle should handle. It must be a pointer.
func (b OpenHandler) NewCommand() interface{} {
	return &OpenCommand{}
}

// Handle ...
func (b OpenHandler) Handle(ctx context.Context, c interface{}) error {
	b.Cmder.Open()

	if err := b.eventBus.Publish(ctx, &es.ValveOpened{
		Time: ptypes.TimestampNow(),
		Id:   uuid.New().String(),
	}); err != nil {
		return err
	}

	return nil
}

// NewOpenHandler OpenHandler constructor
func NewOpenHandler(eb *cqrs.EventBus, cmder Commander) *OpenHandler {
	return &OpenHandler{
		eventBus: eb,
		Cmder:    cmder,
	}
}
