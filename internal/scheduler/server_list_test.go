package scheduler_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/FrancescoIlario/grower/internal/scheduler"
	"github.com/FrancescoIlario/grower/internal/scheduler/memstore"
	vgrpc "github.com/FrancescoIlario/grower/internal/valve/grpc"
	"github.com/FrancescoIlario/grower/internal/valve/mocks"
	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	valvepb "github.com/FrancescoIlario/grower/pkg/valvepb/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var (
	pairsToRead = []*scheduler.Pair{
		{
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
		},
		{
			CreationTime: time.Now(),
			OpenEntryID:  12,
			CloseEntryID: 13,
			OpenSpec:     "16 20 * * *",
			CloseSpec:    "21 20 * * *",
			OpenTime: scheduler.TimePoint{
				Hours:   16,
				Minutes: 20,
			},
			CloseTime: scheduler.TimePoint{
				Hours:   21,
				Minutes: 20,
			},
		},
	}
)

func arrangeList(ctx context.Context, t *testing.T) {
	store = memstore.New()
	for _, pair := range pairsToRead {
		id, err := store.Store(ctx, *pair)
		if err != nil {
			t.Fatalf("error arranging the test: %v", err)
		}
		pair.ID = id.String()
	}

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

func Test_ListSchedules(t *testing.T) {
	ctx := context.Background()
	arrangeList(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	req := &schedulerpb.ListSchedulesRequest{}

	resp, err := client.ListSchedules(ctx, req)
	if err != nil {
		t.Fatalf("error obtaining response: %+v", err)
	}

	if lenExp, lenObt := len(pairsToRead), len(resp.Schedules); lenExp != lenObt {
		t.Fatalf("expected %v schedules, obtained %v", lenExp, lenObt)
	}

	for _, schedule := range resp.Schedules {
		pair := getPairToRead(t, schedule.Id)
		checkPairs(t, pair, schedule)
	}
}

func getPairToRead(t *testing.T, id string) *scheduler.Pair {
	for _, pair := range pairsToRead {
		if idstr := pair.ID; idstr == id {
			return pair
		}
	}

	t.Fatalf("Pair with id %v not found", id)
	return nil
}
