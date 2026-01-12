package domain

import "time"

type Ride struct {
	ID        string
	UserID    string
	DriverID  string
	Status    RideStatus
	Fare      int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *Ride)Transition(next RideStatus)error{
	if !r.Status.CanTransitionTo(next){
		return ErrInvalidRideTransition
	}
	r.Status = next
	return nil
}