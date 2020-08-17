package grpc

import (
	"context"

	vcqrs "github.com/FrancescoIlario/grower/pkg/valvepb/cqrs"
	"github.com/FrancescoIlario/grower/pkg/valvepb/grpc"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type valveServer struct {
	facade cqrs.Facade
}

// NewGrpcServer ...
func NewGrpcServer(commandsPublisher message.Publisher) (grpc.ValveServiceServer, error) {
	facade, err := cqrsFacade(&commandsPublisher)
	if err != nil {
		return nil, err
	}

	return &valveServer{
		facade: *facade,
	}, nil
}

func (v *valveServer) GetStatus(context.Context, *grpc.GetStatusRequest) (*grpc.GetStatusResponse, error) {
	// st := v.Cmder.Status()
	// statuspb := st.toStatusPB()
	// return &valvepb.GetStatusResponse{Status: statuspb}, nil

	// TODO: Create the query and publish it
	return &grpc.GetStatusResponse{}, nil
}

func (v *valveServer) OpenValve(ctx context.Context, req *grpc.OpenValveRequest) (*grpc.OpenValveResponse, error) {
	cmd := &vcqrs.OpenValveCommand{
		Id:           uuid.New().String(),
		CreationTime: ptypes.TimestampNow(),
	}

	if err := v.facade.CommandBus().Send(ctx, cmd); err != nil {
		return nil, status.Errorf(codes.Internal, "error sending open valve command: %v", err)
	}

	return &grpc.OpenValveResponse{}, nil
}

func (v *valveServer) CloseValve(ctx context.Context, req *grpc.CloseValveRequest) (*grpc.CloseValveResponse, error) {
	cmd := &vcqrs.OpenValveCommand{
		Id:           uuid.New().String(),
		CreationTime: ptypes.TimestampNow(),
	}

	cb := v.facade.CommandBus()
	if err := cb.Send(ctx, cmd); err != nil {
		return nil, status.Errorf(codes.Internal, "error sending close valve command: %v", err)
	}

	return &grpc.CloseValveResponse{}, nil
}

func cqrsFacade(commandsPublisher *message.Publisher) (*cqrs.Facade, error) {
	logger := watermill.NewStdLogger(false, false)

	// CQRS is built on messages router. Detailed documentation: https://watermill.io/docs/messages-router/
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}

	cqrsMarshaler := cqrs.ProtobufMarshaler{}

	// cqrs.Facade is facade for Command and Event buses and processors.
	cqrsFacade, err := cqrs.NewFacade(cqrs.FacadeConfig{
		GenerateCommandsTopic: func(commandName string) string {
			// we are using queue RabbitMQ config, so we need to have topic per command type
			return commandName
		},
		CommandsPublisher:     *commandsPublisher,
		Router:                router,
		CommandEventMarshaler: cqrsMarshaler,
		Logger:                logger,
	})
	if err != nil {
		return nil, err
	}

	return cqrsFacade, nil
}
