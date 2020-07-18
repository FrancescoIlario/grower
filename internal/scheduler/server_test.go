package scheduler_test

import (
	"testing"

	"github.com/FrancescoIlario/grower/internal/scheduler"
	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var (
	conn  *grpc.ClientConn
	lis   *bufconn.Listener
	store scheduler.PairStore
)

func checkPairs(t *testing.T, expected *scheduler.Pair, obtained *schedulerpb.Schedule) {
	if expected == nil && obtained != nil {
		t.Fatalf("expected nil, obtained %v", obtained)
	}
	if expected != nil && obtained == nil {
		t.Fatalf("expected %v, obtained nil", expected)
	}

	if exp, obt := expected.OpenTime.Hours, int(obtained.OpenTime.Hours); exp != obt {
		t.Errorf("Open time Hours expected %v, obtained %v", exp, obt)
	}
	if exp, obt := expected.OpenTime.Minutes, int(obtained.OpenTime.Minutes); exp != obt {
		t.Errorf("Open time Minutes expected %v, obtained %v", exp, obt)
	}

	if exp, obt := expected.CloseTime.Hours, int(obtained.CloseTime.Hours); exp != obt {
		t.Errorf("Close time Hours expected %v, obtained %v", exp, obt)
	}
	if exp, obt := expected.CloseTime.Minutes, int(obtained.CloseTime.Minutes); exp != obt {
		t.Errorf("Close time Minutes expected %v, obtained %v", exp, obt)
	}

	if exp, obt := expected.ID, obtained.Id; exp != obt {
		t.Errorf("Close time Minutes expected %v, obtained %v", exp, obt)
	}
	if exp, obt := expected.CreationTime, obtained.CreationTime.AsTime(); exp.Nanosecond() != obt.Nanosecond() {
		t.Errorf("Creation Time obtained %v (%v), expected to be close to %v (%v)", exp, exp.Nanosecond(), obt, obt.Nanosecond())
	}
}
