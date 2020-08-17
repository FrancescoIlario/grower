package proc

import "github.com/FrancescoIlario/grower/pkg/valvepb/shared"

// Status ...
type Status int

const (
	// StatusInvalid the valve is neither open nor closed (invalid tension pair on relays)
	StatusInvalid Status = iota + 1
	// StatusOpening the valve is opening
	StatusOpening
	// StatusOpen the valve is open
	StatusOpen
	// StatusClosing the valve is closing
	StatusClosing
	// StatusClose the valve is closed
	StatusClose
)

func (s *Status) toStatusPB() shared.ValveStatus {
	switch *s {
	case StatusInvalid:
		return shared.ValveStatus_VALVE_INVALID
	case StatusOpening:
		return shared.ValveStatus_VALVE_OPENING
	case StatusOpen:
		return shared.ValveStatus_VALVE_OPEN
	case StatusClosing:
		return shared.ValveStatus_VALVE_CLOSING
	case StatusClose:
		return shared.ValveStatus_VALVE_CLOSE
	default:
		return shared.ValveStatus_VALVE_UNSPECIFIED
	}
}
