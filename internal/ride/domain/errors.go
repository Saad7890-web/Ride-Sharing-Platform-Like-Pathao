package domain

import "errors"

var (
	ErrInvalidRideTransition = errors.New("invalid ride status transition")
)
