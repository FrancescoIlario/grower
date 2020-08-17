package proc

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// Commander ...
type Commander interface {
	Status() Status
	Open()
	Close()
}

type commander struct {
	pulselength      time.Duration
	positiveRelayPin rpio.Pin
	negativeRelayPin rpio.Pin
}

// NewCommander ...
func NewCommander(positiveRelayPin, negativeRelayPin rpio.Pin, pulselength time.Duration) Commander {
	return &commander{
		pulselength:      pulselength,
		positiveRelayPin: positiveRelayPin,
		negativeRelayPin: negativeRelayPin,
	}
}

// IsOpen checks whether the valve is open or not
func (c *commander) Status() Status {
	pv := c.positiveRelayPin.Read()
	nv := c.negativeRelayPin.Read()

	if pv == rpio.High && nv == rpio.Low {
		return StatusOpening
	} else if pv == rpio.High && nv == rpio.High {
		return StatusOpen
	} else if pv == rpio.Low && nv == rpio.High {
		return StatusClosing
	} else {
		return StatusClose
	}
}

// Open opens the valve
// puts tension HIGH to both terminals activating the relays
func (c *commander) Open() {
	c.setLinesValues(rpio.High, rpio.Low)
	time.Sleep(c.pulselength)
	c.setLinesValues(rpio.High, rpio.High)
}

// Close closes the valve
// puts tension LOW to both terminals shutting the relays
func (c *commander) Close() {
	c.setLinesValues(rpio.Low, rpio.High)
	time.Sleep(c.pulselength)
	c.setLinesValues(rpio.Low, rpio.Low)
}

func (c *commander) setLinesValues(p, n rpio.State) {
	c.setLineValue(c.positiveRelayPin, p)
	c.setLineValue(c.negativeRelayPin, n)
}

func (c *commander) setLineValue(l rpio.Pin, s rpio.State) {
	l.Output()
	l.Write(s)
}
