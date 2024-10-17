package main

//go run github.com/pressly/goose/v3/cmd/goose postgres postgres://postgres:postgres@postgresql-delta-eats:5434/delta-eats?sslmode=disable up

//go:generate go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgres://postgres:postgres@localhost:5434/delta-eats?sslmode=disable -path=../internal/ports/database/gen
//go:generate go run github.com/99designs/gqlgen generate
import (
	"hradec/cmd/runner"
	"log"
	"os"
)

func main() {
	err := runner.Serve()
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}
}
