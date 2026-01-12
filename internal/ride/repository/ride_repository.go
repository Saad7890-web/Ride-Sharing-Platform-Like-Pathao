package repository

import (
	"context"

	"ride/internal/ride/domain"
)

type RideRepository interface {
	Create(ctx context.Context, ride *domain.Ride) error
	UpdateStatus(ctx context.Context, rideID string, status domain.RideStatus) error
}
