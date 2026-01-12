package repository

import (
	"context"

	commonRepo "ride/internal/common/repository"
	"ride/internal/ride/domain"
)

type RideRepository interface {
	Create(ctx context.Context, tx commonRepo.Tx ,ride *domain.Ride) error
	UpdateStatus(ctx context.Context, tx commonRepo.Tx, rideID string, status domain.RideStatus) error
	AssignDriver(ctx context.Context, tx commonRepo.Tx, rideID string, driverID string) error
}
