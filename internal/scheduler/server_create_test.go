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

func arrange(ctx context.Context, t *testing.T) {
	store = memstore.New()
	cmder := mocks.NewValveCmder(200 * time.Millisecond)

	s := grpc.NewServer()
	valvepb.RegisterValveServiceServer(s, valve.NewGrpcServer(cmder))

	lis = bufconn.Listen(bufSize)
	var err error
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

func Test_CreateSchedule(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	csr := validCreateScheduleRequest()
	resp, err := client.CreateSchedule(ctx, csr)
	if err != nil {
		t.Fatalf("Create schedule failed: %v", err)
	}

	uid, err := uuid.Parse(resp.Id)
	if err != nil {
		t.Fatalf("invalid uuid returned by the server: %v", resp.Id)
	}

	pairs, err := store.List(ctx)
	if lp := len(pairs); lp != 1 {
		t.Fatalf("expected lp to have only 1 entry, obtained %d", lp)
	}

	p, err := store.Read(ctx, uid)
	if err != nil {
		t.Fatalf("error returned looking for pair with id %s: %v", resp.Id, err)
	}
	if p.ID != uid {
		t.Errorf("expected id %s, obtained %s", uid.String(), p.ID.String())
	}
	if exp, obt := p.CloseTime.Hours, int(csr.GetCloseTime().Hours); exp != obt {
		t.Errorf("close time hours: expected %d, obtained %d", exp, obt)
	}
	if exp, obt := p.CloseTime.Minutes, int(csr.GetCloseTime().Minutes); exp != obt {
		t.Errorf("close time minutes: expected %d, obtained %d", exp, obt)
	}
	if exp, obt := p.OpenTime.Hours, int(csr.GetOpenTime().Hours); exp != obt {
		t.Errorf("open time hours: expected %d, obtained %d", exp, obt)
	}
	if exp, obt := p.OpenTime.Minutes, int(csr.GetOpenTime().Minutes); exp != obt {
		t.Errorf("open time minutes: expected %d, obtained %d", exp, obt)
	}
}

func Test_CreateSchedule_invalid_closetime_hour_less_than_0(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	csr := validCreateScheduleRequest()
	csr.CloseTime.Hours = -1

	if resp, err := client.CreateSchedule(ctx, csr); err == nil {
		t.Fatalf("expected error, obtained response: %+v", resp)
	}
}

func Test_CreateSchedule_invalid_closetime_hour_more_than_23(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	csr := validCreateScheduleRequest()
	csr.CloseTime.Hours = 24

	if resp, err := client.CreateSchedule(ctx, csr); err == nil {
		t.Fatalf("expected error, obtained response: %+v", resp)
	}
}

func Test_CreateSchedule_invalid_closetime_minutes_less_than_0(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	csr := validCreateScheduleRequest()
	csr.CloseTime.Minutes = -1

	if resp, err := client.CreateSchedule(ctx, csr); err == nil {
		t.Fatalf("expected error, obtained response: %+v", resp)
	}
}

func Test_CreateSchedule_invalid_closetime_minutes_more_than_59(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	csr := validCreateScheduleRequest()
	csr.CloseTime.Minutes = 60

	if resp, err := client.CreateSchedule(ctx, csr); err == nil {
		t.Fatalf("expected error, obtained response: %+v", resp)
	}
}

func Test_CreateSchedule_invalid_opentime_hour_less_than_0(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	csr := validCreateScheduleRequest()
	csr.OpenTime.Hours = -1

	if resp, err := client.CreateSchedule(ctx, csr); err == nil {
		t.Fatalf("expected error, obtained response: %+v", resp)
	}
}

func Test_CreateSchedule_invalid_opentime_hour_more_than_23(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	csr := validCreateScheduleRequest()
	csr.OpenTime.Hours = 24

	if resp, err := client.CreateSchedule(ctx, csr); err == nil {
		t.Fatalf("expected error, obtained response: %+v", resp)
	}
}

func Test_CreateSchedule_invalid_opentime_minutes_less_than_0(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	csr := validCreateScheduleRequest()
	csr.OpenTime.Minutes = -1

	if resp, err := client.CreateSchedule(ctx, csr); err == nil {
		t.Fatalf("expected error, obtained response: %+v", resp)
	}
}

func Test_CreateSchedule_invalid_opentime_minutes_more_than_59(t *testing.T) {
	ctx := context.Background()
	arrange(ctx, t)

	client := schedulerpb.NewScheduleServiceClient(conn)
	csr := validCreateScheduleRequest()
	csr.OpenTime.Minutes = 60

	if resp, err := client.CreateSchedule(ctx, csr); err == nil {
		t.Fatalf("expected error, obtained response: %+v", resp)
	}
}

func validCreateScheduleRequest() *schedulerpb.CreateScheduleRequest {
	return &schedulerpb.CreateScheduleRequest{
		OpenTime:  &schedulerpb.TimePoint{Hours: 20, Minutes: 15},
		CloseTime: &schedulerpb.TimePoint{Hours: 20, Minutes: 20},
	}
}
