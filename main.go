package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"todos/internal/task"
)

func main() {
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := task.NewSQLiteRepository(db)
	if err := repo.Migrate(context.Background()); err != nil {
		log.Fatal(err)
	}

	handler := task.NewHandler(repo)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
