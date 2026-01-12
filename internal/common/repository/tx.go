package repository

import "context"


type Tx interface {
	Commit() error
	Rollback() error
}

type TransactionManager interface {
	Begin(ctx context.Context) (Tx, error)
}