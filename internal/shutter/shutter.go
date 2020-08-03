package shutter

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type shutter struct {
	outputPin rpio.Pin
}

//Shutter ...
type Shutter interface {
	shutdown() error
}

//New ...
func New(opin uint) Shutter {
	return &shutter{
		outputPin: rpio.Pin(opin),
	}
}

func (s *shutter) shutdown() error {
	if err := rpio.Open(); err != nil {
		return err
	}
	defer rpio.Close()

	s.outputPin.Output()
	s.outputPin.High()

	return nil
}
