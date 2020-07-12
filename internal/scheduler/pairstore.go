package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

// PairStore ...
type PairStore interface {
	Store(context.Context, Pair) (*uuid.UUID, error)
	Read(context.Context, uuid.UUID) (*Pair, error)
	List(context.Context) ([]Pair, error)
	Delete(context.Context, uuid.UUID) error
}

// Pair ...
type Pair struct {
	ID           uuid.UUID    `bson:"_id" json:"_id"`
	OpenEntryID  cron.EntryID `bson:"open_entry_id" json:"open_entry_id"`
	OpenSpec     string       `bson:"open_spec" json:"open_spec"`
	OpenTime     TimePoint    `bson:"open_time" json:"open_time"`
	CloseEntryID cron.EntryID `bson:"close_entry_id" json:"close_entry_id"`
	CloseSpec    string       `bson:"close_spec" json:"close_spec"`
	CloseTime    TimePoint    `bson:"close_time" json:"close_time"`
	CreationTime time.Time    `bson:"creation_time" json:"creation_time"`
}

//TimePoint ...
type TimePoint struct {
	Hours   int `bson:"hours" json:"hours"`
	Minutes int `bson:"minutes" json:"minutes"`
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
