package repository

import (
	"context"

	"ride/internal/ride/domain"
)

type DriverRepository interface {
	FindAvailable(ctx context.Context, limit int) ([]domain.Driver, error)
	MarkBusy(ctx context.Context, driverID string) error
}
