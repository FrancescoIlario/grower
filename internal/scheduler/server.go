package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	"github.com/FrancescoIlario/grower/pkg/valvepb"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	cc     *cron.Cron
	client valvepb.ValveServiceClient
	mctx   context.Context
	store  PairStore
}

// NewServer ...
func NewServer(valveCmdrHost string, store PairStore) schedulerpb.ScheduleServiceServer {
	conn, err := grpc.Dial(valveCmdrHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	return &server{
		cc:     cron.New(),
		client: valvepb.NewValveServiceClient(conn),
		mctx:   context.Background(),
		store:  store,
	}
}

func (s *server) ListSchedules(ctx context.Context, _ *schedulerpb.ListSchedulesRequest) (*schedulerpb.ListSchedulesResponse, error) {
	pp, err := s.store.List(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("error reading list of stored pairs: %v", err))
	}

	sl := make([]*schedulerpb.Schedule, len(pp))
	for i, p := range pp {
		sl[i] = p.toSchedule()
	}
	return &schedulerpb.ListSchedulesResponse{Schedules: sl}, nil
}

func (s *server) GetSchedule(ctx context.Context, gsr *schedulerpb.GetScheduleRequest) (*schedulerpb.Schedule, error) {
	uid, err := uuid.Parse(gsr.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id %s is not a valid UUID: %v", gsr.Id, err)
	}
	p, err := s.store.Read(ctx, uid)
	if err != nil {
		switch err {
		case ErrNotFound:
			return nil, status.Errorf(codes.NotFound, "id %s not found", gsr.Id)
		default:
			return nil, status.Errorf(codes.Internal, "error retrieving entry with id %s: %v", gsr.Id, err)
		}
	}
	return p.toSchedule(), nil
}

func (s *server) CreateSchedule(ctx context.Context, req *schedulerpb.CreateScheduleRequest) (*schedulerpb.Schedule, error) {
	if err := req.OpenTime.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid OpenTime: %v", err))
	}
	if err := req.CloseTime.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid CloseTime: %v", err))
	}

	ospec, cspec := toCron(req.OpenTime), toCron(req.CloseTime)
	otentry, err := s.cc.AddFunc(ospec, func() { s.open() })
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("could not create entry for OpenTime: %v", req.OpenTime))
	}
	ctentry, err := s.cc.AddFunc(cspec, func() { s.close() })
	if err != nil {
		s.cc.Remove(otentry)
		return nil, status.Error(codes.Internal, fmt.Sprintf("could not create entry for CloseTime: %v", req.CloseTime))
	}

	p := Pair{
		OpenEntryID:  otentry,
		OpenSpec:     ospec,
		CloseEntryID: ctentry,
		CloseSpec:    cspec,
		CreationTime: time.Now(),
	}

	id, err := s.store.Store(ctx, p)
	if err != nil {
		s.cc.Remove(ctentry)
		s.cc.Remove(otentry)
		return nil, status.Error(codes.Internal, fmt.Sprintf("could not store pair in pairstore: %v", err))
	}

	tp, _ := ptypes.TimestampProto(p.CreationTime)
	return &schedulerpb.Schedule{
		Id:           id.String(),
		OpenTime:     req.OpenTime,
		CloseTime:    req.CloseTime,
		CreationTime: tp,
	}, nil
}

func (s *server) DeleteSchedule(ctx context.Context, dsr *schedulerpb.DeleteScheduleRequest) (*empty.Empty, error) {
	uid, err := uuid.Parse(dsr.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id %s is not a valid UUID: %v", dsr.Id, err)
	}

	if err := s.store.Delete(ctx, uid); err != nil {
		switch err {
		case ErrNotFound:
			return nil, status.Errorf(codes.NotFound, "id %s not found", dsr.Id)
		default:
			return nil, status.Error(codes.Internal, fmt.Sprintf("could not delete pair with id %s: %v", dsr.Id, err))
		}
	}
	return &emptypb.Empty{}, nil
}

func (s *server) open() {
	logrus.Debugf("opening %v", time.Now())
	cont, cancel := context.WithTimeout(s.mctx, 1*time.Minute)
	defer cancel()

	s.client.OpenValve(cont, &valvepb.OpenValveRequest{})
	logrus.Debugf("opened %v", time.Now())
}

func (s *server) close() {
	logrus.Debugf("closing %v", time.Now())
	cont, cancel := context.WithTimeout(s.mctx, 1*time.Minute)
	defer cancel()

	s.client.CloseValve(cont, &valvepb.CloseValveRequest{})
	logrus.Debugf("closeded %v", time.Now())
}

func toCron(t *schedulerpb.TimePoint) string {
	return fmt.Sprintf("%d %d * * *", t.GetMinutes(), t.GetHours())
}
