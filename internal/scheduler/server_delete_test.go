package scheduler_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/FrancescoIlario/grower/internal/mocks"
	"github.com/FrancescoIlario/grower/internal/scheduler"
	"github.com/FrancescoIlario/grower/internal/scheduler/memstore"
	"github.com/FrancescoIlario/grower/internal/valve"
	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	"github.com/FrancescoIlario/grower/pkg/valvepb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var (
	pairToDeleteID *uuid.UUID
	pairToDelete   = scheduler.Pair{
		CreationTime: time.Now(),
		OpenEntryID:  10,
		CloseEntryID: 11,
		OpenSpec:     "15 20 * * *",
		CloseSpec:    "20 20 * * *",
		OpenTime: scheduler.TimePoint{
			Hours:   20,
			Minutes: 15,
		},
		CloseTime: scheduler.TimePoint{
			Hours:   20,
			Minutes: 20,
		},
	}
)

func arrangeDelete(ctx context.Context, t *testing.T) {
	store = memstore.New()
	id, err := store.Store(ctx, pairToDelete)
	if err != nil {
		t.Fatalf("error arranging the test: %v", err)
	}
	pairToDeleteID = id
	pairToDelete.ID = id.String()

	cmder := mocks.NewValveCmder(200 * time.Millisecond)
	s := grpc.NewServer()
	valvepb.RegisterValveServiceServer(s, valve.NewGrpcServer(cmder))

	lis = bufconn.Listen(bufSize)
	conn, err = grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("Failed to dial bufnet: %w", err))
	}

	valvecli := valvepb.NewValveServiceClient(conn)
	schedulerpb.RegisterScheduleServiceServer(s, scheduler.NewServer(store, valvecli))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func Test_DeleteSchedule(t *testing.T) {
	ctx := context.Background()
	arrangeDelete(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	req := &schedulerpb.DeleteScheduleRequest{Id: pairToDelete.ID}

	if _, err := client.DeleteSchedule(ctx, req); err != nil {
		t.Fatalf("error deleting schedule: %+v", err)
	}

	_, err := store.Read(ctx, *pairToDeleteID)
	if err == nil || err != scheduler.ErrNotFound {
		t.Fatalf("expected ErrNotFound, obtained %v", err)
	}
}
