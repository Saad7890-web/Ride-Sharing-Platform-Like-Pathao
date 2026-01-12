package repository

import (
	"context"

	commonRepo "ride/internal/common/repository"
	"ride/internal/ride/domain"
)

type DriverRepository interface {
	FindAvailable(ctx context.Context, limit int) ([]domain.Driver, error)
	MarkBusy(ctx context.Context,tx commonRepo.Tx, driverID string) error
}
