package main

import (
	"database/sql"
	"log"
	
	"github.com/katatrina/SWD392_NET1701_GroupIntern/api"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern/db/sqlc"
	
	_ "github.com/lib/pq"
)

const (
	dbSource = "postgres://root:secret@localhost/dental_clinic?sslmode=disable"
)

//	@title			Dental Clinic API
//	@version		1.0.0
//	@description	API documentation for the Dental Clinic application.

//	@contact.name	Châu Vĩnh Phước
//	@contact.email	cvphuoc2014@gmail.com

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.apikey	accessToken
// @in							header
// @name						Authorization
// @description				JWT Authorization header using the Bearer scheme.
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
