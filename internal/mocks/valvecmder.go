package mocks

import (
	"time"

	"github.com/FrancescoIlario/grower/internal/valve/proc"
)

// ValveCmder ...
type valveCmder struct {
	status proc.Status
	Delay  time.Duration
}

// NewValveCmder ...
func NewValveCmder(delay time.Duration) proc.Commander {
	return &valveCmder{
		Delay: delay,
	}
}

// Status ...
func (v *valveCmder) Status() proc.Status {
	return v.status
}

// Open ...
func (v *valveCmder) Open() {
	v.status = proc.StatusOpening
	time.Sleep(v.Delay)
	v.status = proc.StatusOpen
}

// Close ...
func (v *valveCmder) Close() {
	v.status = proc.StatusClosing
	time.Sleep(v.Delay)
	v.status = proc.StatusClose
}
