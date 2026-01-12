package database

import (
	"context"
	"database/sql"
	"ride/internal/common/repository"
)

type postgresTx struct {
	tx *sql.Tx
}

func (p *postgresTx) Commit()error {
	return p.tx.Commit()
}

func (p *postgresTx) Rollback()error{
	return p.tx.Rollback()
}

type PostgresTxManager struct {
	db *sql.DB
}

func NewPostgresTxManager(db *sql.DB)repository.TransactionManager{
	return &PostgresTxManager{db: db}
}

func (m *PostgresTxManager) Begin(ctx context.Context) (repository.Tx, error){
	tx, err := m.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	return &postgresTx{tx: tx}, nil
}