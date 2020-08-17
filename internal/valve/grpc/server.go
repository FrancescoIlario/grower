package grpc

import (
	"context"
	"fmt"

	vcqrs "github.com/FrancescoIlario/grower/pkg/valvepb/cqrs"
	"github.com/FrancescoIlario/grower/pkg/valvepb/grpc"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type valveServer struct {
	commandBus *cqrs.CommandBus
}

// NewGrpcServer ...
func NewGrpcServer(commandsPublisher message.Publisher) (grpc.ValveServiceServer, error) {
	cqrsMarshaler := cqrs.ProtobufMarshaler{}

	commandBus, err := cqrs.NewCommandBus(
		commandsPublisher,
		func(commandName string) string {
			// we are using queue RabbitMQ config, so we need to have topic per command type
			return commandName
		},
		cqrsMarshaler,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create command bus: %w", err)
	}
	return &valveServer{commandBus: commandBus}, nil
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

	if err := v.commandBus.Send(ctx, cmd); err != nil {
		return nil, status.Errorf(codes.Internal, "error sending open valve command: %v", err)
	}

	return &grpc.OpenValveResponse{}, nil
}

func (v *valveServer) CloseValve(ctx context.Context, req *grpc.CloseValveRequest) (*grpc.CloseValveResponse, error) {
	cmd := &vcqrs.OpenValveCommand{
		Id:           uuid.New().String(),
		CreationTime: ptypes.TimestampNow(),
	}

	if err := v.commandBus.Send(ctx, cmd); err != nil {
		return nil, status.Errorf(codes.Internal, "error sending close valve command: %v", err)
	}

	return &grpc.CloseValveResponse{}, nil
}
