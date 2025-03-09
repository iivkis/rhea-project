package ports

import (
	"context"
)

type ApiState struct {
	UserService IUserService
	PgExec      PgExecutor
}

// db adapters
type Row interface {
	Scan(dest ...any) error
}

type Executor interface {
	QueryRow(ctx context.Context, query string, args ...any) Row
	Begin(ctx context.Context) (PgExecutorTx, error)
}

type PgExecutor interface {
	Executor
	isPgExecutor()
}

type PgExecutorTx interface {
	PgExecutor
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
