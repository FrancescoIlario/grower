package mongostore_test

import (
	"context"
	"testing"
	"time"

	"github.com/FrancescoIlario/grower/internal/scheduler"
)

func TestCreateGet_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	ctx := context.TODO()
	store, pairToCreate := generateNewStoreOrFatal(t), getPairToCreate()
	uid, err := store.Store(ctx, *pairToCreate)
	if err != nil {
		t.Fatalf("error storing pair: %v", err)
	}
	pairToCreate.ID = uid.String()

	pairRead, err := store.Read(ctx, *uid)
	if err != nil {
		t.Fatalf("error reading pair: %v", err)
	}

	checkPairs(t, pairToCreate, pairRead)
}

func TestCreateList_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	ctx := context.TODO()
	store, pairToCreate := generateNewStoreOrFatal(t), getPairToCreate()
	uid, err := store.Store(ctx, *pairToCreate)
	if err != nil {
		t.Fatalf("error storing pair: %v", err)
	}
	pairToCreate.ID = uid.String()

	pairsRead, err := store.List(ctx)
	if err != nil {
		t.Fatalf("error reading pair: %v", err)
	}
	if lpr := len(pairsRead); lpr != 1 {
		t.Fatalf("expected 1 pairs, obtained %v", lpr)
	}

	checkPairs(t, pairToCreate, &pairsRead[0])
}

func TestCreateDelete_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	ctx := context.TODO()
	store, pairToCreate := generateNewStoreOrFatal(t), getPairToCreate()
	uid, err := store.Store(ctx, *pairToCreate)
	if err != nil {
		t.Fatalf("error storing pair: %v", err)
	}
	pairToCreate.ID = uid.String()

	pairRead, err := store.Read(ctx, *uid)
	if err != nil {
		t.Fatalf("error reading pair: %v", err)
	}

	checkPairs(t, pairToCreate, pairRead)

	if err := store.Delete(ctx, *uid); err != nil {
		t.Fatalf("error deleting pair: %v", err)
	}

	if _, err := store.Read(ctx, *uid); err == nil {
		t.Fatalf("expected error after deletion not returned")
	}
}

func getPairToCreate() *scheduler.Pair {
	return &scheduler.Pair{
		OpenEntryID:  10,
		CloseEntryID: 11,
		CloseTime: scheduler.TimePoint{
			Hours:   20,
			Minutes: 15,
		},
		OpenTime: scheduler.TimePoint{
			Hours:   20,
			Minutes: 20,
		},
		CreationTime: time.Now(),
		CloseSpec:    "15 20 * * *",
		OpenSpec:     "20 20 * * *",
	}
}

func checkPairs(t *testing.T, expected, obtained *scheduler.Pair) {
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

	if exp, obt := expected.ID, obtained.ID; exp != obt {
		t.Errorf("Close time Minutes expected %v, obtained %v", exp, obt)
	}
	if exp, obt := expected.CreationTime, obtained.CreationTime; exp.Unix() != obt.Unix() {
		t.Errorf("Creation Time obtained %v (%v), expected to be close to %v (%v)", exp, exp.Unix(), obt, obt.Unix())
	}
}
