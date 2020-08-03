package shutter

import (
	"context"

	"github.com/FrancescoIlario/grower/pkg/shutterpb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type shutterServer struct {
	shutter Shutter
}

// NewGrpcServer ...
func NewGrpcServer(outputPin uint) shutterpb.ShutterServiceServer {
	return &shutterServer{
		shutter: New(outputPin),
	}
}

func (ss *shutterServer) Shut(context.Context, *shutterpb.ShutRequest) (*empty.Empty, error) {
	if err := ss.shutter.shutdown(); err != nil {
		return nil, status.Errorf(codes.Internal, "error shutting down: %v", err)
	}
	return &emptypb.Empty{}, nil
}
