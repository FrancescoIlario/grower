package valve

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
	positiveRelayPin rpio.Pin
	negativeRelayPin rpio.Pin
}

// NewCommander ...
func NewCommander(positiveRelayPin, negativeRelayPin rpio.Pin) Commander {
	return &commander{
		positiveRelayPin: positiveRelayPin,
		negativeRelayPin: negativeRelayPin,
	}
}

// IsOpen checks whether the valve is open or not
func (c *commander) Status() Status {
	pv := c.positiveRelayPin.Read()
	nv := c.negativeRelayPin.Read()

	if pv == rpio.High && nv == rpio.High {
		return StatusOpening
	} else if pv == rpio.High && nv == rpio.Low {
		return StatusOpen
	} else if pv == rpio.Low && nv == rpio.Low {
		return StatusClosing
	} else {
		return StatusClose
	}
}

// Open opens the valve
// puts tension HIGH to both terminals activating the relays
func (c *commander) Open() {
	c.setLinesValues(rpio.High, rpio.High)
	time.Sleep(100 * time.Millisecond)
	c.setLinesValues(rpio.High, rpio.Low)
}

// Close closes the valve
// puts tension LOW to both terminals shutting the relays
func (c *commander) Close() {
	c.setLinesValues(rpio.Low, rpio.Low)
	time.Sleep(100 * time.Millisecond)
	c.setLinesValues(rpio.Low, rpio.High)
}

func (c *commander) setLinesValues(p, n rpio.State) {
	c.setLineValue(c.positiveRelayPin, p)
	c.setLineValue(c.negativeRelayPin, n)
}

func (c *commander) setLineValue(l rpio.Pin, s rpio.State) {
	l.Output()
	l.Write(s)
}
