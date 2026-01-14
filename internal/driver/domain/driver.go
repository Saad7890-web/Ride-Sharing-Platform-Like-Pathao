package domain

import "time"

type DriverStatus string

const (
	DriverPending   DriverStatus = "PENDING"
	DriverApproved  DriverStatus = "APPROVED"
	DriverRejected  DriverStatus = "REJECTED"
	DriverSuspended DriverStatus = "SUSPENDED"
)

type Driver struct {
	ID        string
	UserID    string

	Status    DriverStatus
	IsOnline  bool
	IsBusy    bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
