package mocks

import (
	"time"

	"github.com/FrancescoIlario/grower/internal/valve"
)

// ValveCmder ...
type valveCmder struct {
	status valve.Status
	Delay  time.Duration
}

// NewValveCmder ...
func NewValveCmder(delay time.Duration) valve.Commander {
	return &valveCmder{
		Delay: delay,
	}
}

// Status ...
func (v *valveCmder) Status() valve.Status {
	return v.status
}

// Open ...
func (v *valveCmder) Open() {
	v.status = valve.StatusOpening
	time.Sleep(v.Delay)
	v.status = valve.StatusOpen
}

// Close ...
func (v *valveCmder) Close() {
	v.status = valve.StatusClosing
	time.Sleep(v.Delay)
	v.status = valve.StatusClose
}
