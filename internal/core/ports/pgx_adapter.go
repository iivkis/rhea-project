package ports

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PgxPoolAdapter - адаптер для пула
type PgxPoolAdapter struct {
	pool *pgxpool.Pool
}

func NewPgxPoolAdapter(pool *pgxpool.Pool) PgExecutor {
	return &PgxPoolAdapter{pool}
}

func (p *PgxPoolAdapter) QueryRow(ctx context.Context, query string, args ...any) Row {
	return p.pool.QueryRow(ctx, query, args)
}

func (p *PgxPoolAdapter) Begin(ctx context.Context) (PgExecutorTx, error) {
	tx, err := p.pool.Begin(ctx)
	return NewPgxTxAdapter(tx), err
}

func (p *PgxPoolAdapter) isPgExecutor() {}

// PgxTxAdapter - адаптер для tx.Begin()
type PgxTxAdapter struct {
	tx pgx.Tx
}

func NewPgxTxAdapter(tx pgx.Tx) PgExecutorTx {
	return &PgxTxAdapter{tx}
}

func (p *PgxTxAdapter) QueryRow(ctx context.Context, query string, args ...any) Row {
	return p.tx.QueryRow(ctx, query, args)
}

func (p *PgxTxAdapter) Commit(ctx context.Context) error {
	return p.tx.Commit(ctx)
}

func (p *PgxTxAdapter) Rollback(ctx context.Context) error {
	return p.tx.Rollback(ctx)
}

func (p *PgxTxAdapter) Begin(ctx context.Context) (PgExecutorTx, error) {
	tx, err := p.tx.Begin(ctx)
	return NewPgxTxAdapter(tx), err
}

func (p *PgxTxAdapter) isPgExecutor() {}
