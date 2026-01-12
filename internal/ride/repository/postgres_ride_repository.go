package repository

import (
	"context"
	"database/sql"

	commonRepo "ride/internal/common/repository"
	"ride/internal/ride/domain"
)

type PostgresRideRepository struct {
	db *sql.DB
}

func NewPostgresRideRepository(db *sql.DB) *PostgresRideRepository {
	return &PostgresRideRepository{db: db}
}

func (r *PostgresRideRepository) Create(
	ctx context.Context,
	tx commonRepo.Tx,
	ride *domain.Ride,
) error {

	return tx.ExecContext(ctx, `
		INSERT INTO rides (id, passenger_id, status, created_at)
		VALUES ($1, $2, $3, NOW())
	`, ride.ID, ride.UserID, ride.Status)
}

func (r *PostgresRideRepository) UpdateStatus(
	ctx context.Context,
	tx commonRepo.Tx,
	rideID string,
	status domain.RideStatus,
) error {

	return tx.ExecContext(ctx, `
		UPDATE rides SET status=$1 WHERE id=$2
	`, status, rideID)
}

func (r *PostgresRideRepository) AssignDriver(
	ctx context.Context,
	tx commonRepo.Tx,
	rideID string,
	driverID string,
) error {

	return tx.ExecContext(ctx, `
		UPDATE rides SET driver_id=$1 WHERE id=$2
	`, driverID, rideID)
}
