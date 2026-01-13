package database

import (
	"context"

	commonRepo "ride/internal/common/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresTx struct {
	tx pgx.Tx
}


func (p *postgresTx) ExecContext(ctx context.Context, query string, args ...any) error {
	_, err := p.tx.Exec(ctx, query, args...)
	return err
}


func (p *postgresTx) Commit() error {
	return p.tx.Commit(context.Background())
}


func (p *postgresTx) Rollback() error {
	return p.tx.Rollback(context.Background())
}


type PostgresTxManager struct {
	db *pgxpool.Pool
}


func NewPostgresTxManager(db *pgxpool.Pool) commonRepo.TransactionManager {
	return &PostgresTxManager{db: db}
}


func (m *PostgresTxManager) Begin(ctx context.Context) (commonRepo.Tx, error) {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &postgresTx{tx: tx}, nil
}
