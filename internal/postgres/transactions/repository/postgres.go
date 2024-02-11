package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	*pgxpool.Pool
}

func NewPostgres(dsn string) (*Postgres, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return &Postgres{pool}, nil
}

func (p *Postgres) UseTransaction(ctx context.Context, txFunc func(pgx.Tx) error) error {
	tx, err := p.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Rollback(ctx)
		if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			log.Println(err)
			return
		}
	}()

	err = txFunc(tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (p *Postgres) GetTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := p.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (p *Postgres) CommitTransaction(ctx context.Context, tx pgx.Tx) error {
	err := tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) RollbackTransaction(ctx context.Context, tx pgx.Tx) error {
	err := tx.Rollback(ctx)
	if err != nil {
		return err
	}

	return nil
}
