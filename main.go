package main

import (
	"database/sql"
	"log"

	"github.com/katatrina/SWD392/api"
	db "github.com/katatrina/SWD392/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbSource = "postgres://root:secret@localhost/dental_clinic?sslmode=disable"
)

func main() {
	// Initialize database connection pool
	dbConn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	// Initialize database store
	store := db.NewStore(dbConn)

	// Initialize our HTTP server
	server := api.NewServer(store)

	// Start our HTTP server
	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
