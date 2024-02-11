package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type repo struct {
	repo *Postgres
}

func New(dsn string) (*repo, error) {
	r, err := NewPostgres(dsn)
	if err != nil {
		return nil, err
	}

	return &repo{r}, nil
}

func (r *repo) Create(ctx context.Context, username string, userDescription string) error {
	err := r.repo.UseTransaction(ctx, func(tx pgx.Tx) error {
		tag, err := tx.Exec(ctx, "INSERT INTO users VALUES ($1, $2)", username, userDescription)
		if err != nil {
			return err
		}

		if tag.RowsAffected() == 0 {
			return errors.New("no rows were added")
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
