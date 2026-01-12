package domain

type RideStatus string

const (
	RideRequested RideStatus = "REQUESTED"
	RideAccepted  RideStatus = "ACCEPTED"
	RideStarted   RideStatus = "STARTED"
	RideCompleted RideStatus = "COMPLETED"
	RideCancelled RideStatus = "CANCELLED"
)

func (s RideStatus) CanTransitionTo(next RideStatus)bool {
	switch s {
	case RideRequested:
		return next == RideAccepted || next == RideCancelled

	case RideAccepted:
		return next == RideStarted || next == RideCancelled

	case RideStarted:
		return next == RideCompleted

	default:
		return false
	}
}