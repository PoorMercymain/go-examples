package main

import (
	"log"

	"github.com/PoorMercymain/go-examples/internal/postgres/repository"
)

func main() {
	err := repository.ApplyMigrations("file://migrations", "postgres://go-examples:go-examples@localhost:5432/go-examples?sslmode=disable")
	if err != nil {
		log.Println(err)
	}
}
