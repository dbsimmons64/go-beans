package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dbsimmons64/go-beans/database"
	"github.com/dbsimmons64/go-beans/repos"
	"github.com/dbsimmons64/go-beans/services"
	"github.com/dbsimmons64/go-beans/sessions"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	templateCache      TemplateCache
	transactionService services.TransactionServiceImpl
	sessionStore       sessions.SessionStore
}

func main() {

	// Cache all the templates
	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Create and Open the database
	db, err := sql.Open("sqlite3", "beans.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	database.Create_txn_table(db)

	// Ensure we close any database connections
	defer db.Close()

	repo := repos.TransactionRepositoryDB{DB: db}
	service := services.TransactionServiceImpl{Repo: repo}
	sessionStore := sessions.NewInMemorySessionStore()

	app := app{
		templateCache:      templateCache,
		transactionService: service,
		sessionStore:       sessionStore,
	}

	// Configure the server
	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	// and start the server
	log.Println("Listening on port 8080")
	srv.ListenAndServe()
}
