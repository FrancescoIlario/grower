package proc

import (
	"context"

	vcqrs "github.com/FrancescoIlario/grower/pkg/valvepb/cqrs"
	"github.com/FrancescoIlario/grower/pkg/valvepb/es"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
)

// CloseHandler is a command handler, which handles CloseCommand and emits ValveCloseed.
//
// In CQRS, one command must be handled by only one handler.
// When another handler with this command is added to command processor, error will be retuerned.
type CloseHandler struct {
	eventBus *cqrs.EventBus
	Cmder    Commander
}

// HandlerName ...
func (b CloseHandler) HandlerName() string {
	return "CloseHandler"
}

// NewCommand returns type of command which this handle should handle. It must be a pointer.
func (b CloseHandler) NewCommand() interface{} {
	return &vcqrs.CloseValveCommand{}
}

// Handle ...
func (b CloseHandler) Handle(ctx context.Context, c interface{}) error {
	b.Cmder.Close()

	if err := b.eventBus.Publish(ctx, &es.ValveClosed{
		Id: uuid.New().String(),
	}); err != nil {
		return err
	}

	return nil
}

// NewCloseHandler CloseHandler constructor
func NewCloseHandler(eb *cqrs.EventBus, cmder Commander) *CloseHandler {
	return &CloseHandler{
		eventBus: eb,
		Cmder:    cmder,
	}
}
