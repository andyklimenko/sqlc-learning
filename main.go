package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/andyklimenko/sqlc-learning/migrate"
	"github.com/andyklimenko/sqlc-learning/tutorial"
	_ "github.com/lib/pq"
)

func main() {
	const dsn = "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err := migrate.UP(dsn); err != nil {
		panic(err)
	}

	queries := tutorial.New(db)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fetch authors before insert: %v\n", authors)

	authorParams := tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	}
	insertedAuthor, err := queries.CreateAuthor(ctx, authorParams)
	if err != nil {
		panic(err)
	}
	log.Printf("created author: %v\n", insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		panic(err)
	}
	log.Printf("fetched author: %v\n", fetchedAuthor)
}
