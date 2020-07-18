package memstore

import (
	"context"
	"fmt"
	"reflect"

	"github.com/FrancescoIlario/grower/internal/scheduler"
	"github.com/google/uuid"
)

// New ...
func New() scheduler.PairStore {
	return &memstore{
		pairs: make(map[uuid.UUID]scheduler.Pair),
	}
}

type memstore struct {
	pairs map[uuid.UUID]scheduler.Pair
}

func (m *memstore) Store(_ context.Context, p scheduler.Pair) (*uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("error creating a new uuid: %w", err)
	}
	p.ID = id
	m.pairs[id] = p
	return &id, nil
}

func (m *memstore) Read(_ context.Context, id uuid.UUID) (*scheduler.Pair, error) {
	if p, ok := m.pairs[id]; ok {
		return &p, nil
	}
	return nil, scheduler.ErrNotFound
}

func (m *memstore) List(_ context.Context) ([]scheduler.Pair, error) {
	// make a copy for security
	pairs := make([]scheduler.Pair, len(m.pairs))
	cpairs := make([]scheduler.Pair, len(m.pairs))
	i := 0
	for _, pair := range m.pairs {
		cpairs[i] = pair
		i++
	}
	reflect.Copy(reflect.ValueOf(pairs), reflect.ValueOf(cpairs))
	return pairs, nil
}

func (m *memstore) Delete(_ context.Context, id uuid.UUID) error {
	if _, ok := m.pairs[id]; ok {
		delete(m.pairs, id)
		return nil
	}

	return scheduler.ErrNotFound
}
