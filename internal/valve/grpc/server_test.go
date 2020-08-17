package grpc_test

import (
	"context"
	"log"
	"net"
	"testing"

	vgrpc "github.com/FrancescoIlario/grower/internal/valve/grpc"
	"github.com/FrancescoIlario/grower/internal/valve/mocks"
	valvepb "github.com/FrancescoIlario/grower/pkg/valvepb/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var (
	conn *grpc.ClientConn
	lis  *bufconn.Listener
)

func arrange(ctx context.Context, t *testing.T) {
	publisher := mocks.DefaultPublisher()
	subscriber := mocks.DefaultSubscriber()

	s := grpc.NewServer()
	vlvsrv, err := vgrpc.NewGrpcServer(publisher, subscriber)
	if err != nil {
		log.Fatalf("Failed to create grpc server: %v", err)
	}
	valvepb.RegisterValveServiceServer(s, vlvsrv)

	lis = bufconn.Listen(bufSize)
	conn, err = grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func Test_OpenValve(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := valvepb.NewValveServiceClient(conn)
	_, err := client.OpenValve(ctx, &valvepb.OpenValveRequest{})
	if err != nil {
		t.Fatalf("error invoking endpoint: %v", err)
	}
}

func Test_CloseValve(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := valvepb.NewValveServiceClient(conn)
	_, err := client.CloseValve(ctx, &valvepb.CloseValveRequest{})
	if err != nil {
		t.Fatalf("error invoking endpoint: %v", err)
	}
}
