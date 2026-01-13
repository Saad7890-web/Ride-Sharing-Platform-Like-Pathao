package repository

import (
	"context"

	commonRepo "ride/internal/common/repository"
	"ride/internal/ride/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)


type PostgresDriverRepository struct {
	db *pgxpool.Pool
}


func NewPostgresDriverRepository(db *pgxpool.Pool) *PostgresDriverRepository {
	return &PostgresDriverRepository{db: db}
}


func (r *PostgresDriverRepository) FindAvailable(
	ctx context.Context,
	limit int,
) ([]domain.Driver, error) {

	rows, err := r.db.Query(ctx, `
		SELECT id, is_busy, is_online
		FROM drivers
		WHERE is_busy = false AND is_online = true
		ORDER BY id
		LIMIT $1
		FOR UPDATE SKIP LOCKED
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []domain.Driver
	for rows.Next() {
		var d domain.Driver
		if err := rows.Scan(&d.ID, &d.IsBusy, &d.IsOnline); err != nil {
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
		UPDATE drivers SET is_busy = true WHERE id = $1
	`, driverID)
}
