package repository

import (
	"context"
	"database/sql"

	commonRepo "ride/internal/common/repository"
	"ride/internal/ride/domain"
)

type PostgresDriverRepository struct {
	db *sql.DB
}

func NewPostgresDriverRepository(db *sql.DB) *PostgresDriverRepository {
	return &PostgresDriverRepository{db: db}
}

func (r *PostgresDriverRepository) FindAvailable(
	ctx context.Context,
	limit int,
) ([]domain.Driver, error) {

	rows, err := r.db.QueryContext(ctx, `
		SELECT id, is_busy
		FROM drivers
		WHERE is_busy = false
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []domain.Driver

	for rows.Next() {
		var d domain.Driver
		if err := rows.Scan(&d.ID, &d.IsBusy); err != nil {
			return nil, err
		}
		drivers = append(drivers, d)
	}

	return drivers, nil
}

func (r *PostgresDriverRepository) MarkBusy(
	ctx context.Context,
	tx commonRepo.Tx,
	driverID string,
) error {

	return tx.ExecContext(ctx, `
		UPDATE drivers SET is_busy = true WHERE id=$1
	`, driverID)
}
