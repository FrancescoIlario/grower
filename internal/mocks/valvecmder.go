package mocks

import (
	"time"

	"github.com/FrancescoIlario/grower/internal/valve/proc"
)

type valveCmder struct {
	status       proc.Status
	Delay        time.Duration
	openCounter  int
	closeCounter int
}

// Commander ...
type Commander interface {
	proc.Commander

	OpenInvokation() int
	CloseInvokation() int
}

// NewValveCmder ...
func NewValveCmder(delay time.Duration) Commander {
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
	v.openCounter++
	v.status = proc.StatusOpening
	time.Sleep(v.Delay)
	v.status = proc.StatusOpen
}

// Close ...
func (v *valveCmder) Close() {
	v.closeCounter++
	v.status = proc.StatusClosing
	time.Sleep(v.Delay)
	v.status = proc.StatusClose
}

func (v *valveCmder) OpenInvokation() int {
	return v.openCounter
}

func (v *valveCmder) CloseInvokation() int {
	return v.closeCounter
}
