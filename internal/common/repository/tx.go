package repository

import (
	"context"
)


type Tx interface {
	ExecContext(ctx context.Context, query string, args ...any) error
	Commit() error
	Rollback() error
}

type TransactionManager interface {
	Begin(ctx context.Context) (Tx, error)
}