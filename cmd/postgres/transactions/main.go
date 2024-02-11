package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"

	migrations "github.com/PoorMercymain/go-examples/internal/postgres/migrations/repository"
	transactions "github.com/PoorMercymain/go-examples/internal/postgres/transactions/repository"
)

func main() {
	err := migrations.ApplyMigrations("file://migrations", "postgres://go-examples:go-examples@localhost:5432/go-examples?sslmode=disable")
	if err != nil {
		log.Println(err)
		return
	}

	p, err := transactions.New("postgres://go-examples:go-examples@localhost:5432/go-examples?sslmode=disable")
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	go func() {
		if ctx.Err() != nil && errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println(ctx.Err())
			cancel()
		}
	}()

	err = p.UseTransaction(ctx, func(tx pgx.Tx) error {
		tx.Exec(ctx, "INSERT INTO users VALUES ($1, $2)", "ab", "no")
		return nil
	})

	if err != nil {
		log.Println(err)
		return
	}
}
