package scheduler

import (
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

// PairStore ...
type PairStore interface {
	Store(Pair) (*uuid.UUID, error)
	Read(uuid.UUID) (*Pair, error)
	List() ([]Pair, error)
	Delete(uuid.UUID) (*Pair, error)
}

// Pair ...
type Pair struct {
	ID           uuid.UUID
	OpenEntryID  cron.EntryID
	OpenSpec     string
	OpenTime     TimePoint
	CloseEntryID cron.EntryID
	CloseSpec    string
	CloseTime    TimePoint
	CreationTime time.Time
}

//TimePoint ...
type TimePoint struct {
	Hours   int
	Minutes int
}

func (p *Pair) toSchedule() *schedulerpb.Schedule {
	tp, _ := ptypes.TimestampProto(p.CreationTime)
	return &schedulerpb.Schedule{
		Id: p.ID.String(),
		OpenTime: &schedulerpb.TimePoint{
			Hours:   int32(p.OpenTime.Hours),
			Minutes: int32(p.OpenTime.Minutes),
		},
		CloseTime: &schedulerpb.TimePoint{
			Hours:   int32(p.CloseTime.Hours),
			Minutes: int32(p.CloseTime.Minutes),
		},
		CreationTime: tp,
	}
}

// ErrNotFound ...
var ErrNotFound = fmt.Errorf("Not found")
