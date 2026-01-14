package domain

import "time"

type DriverLocation struct {
	DriverID  string

	Latitude  float64
	Longitude float64

	UpdatedAt time.Time
}
