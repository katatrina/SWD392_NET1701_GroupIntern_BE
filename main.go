package main

import (
	"database/sql"
	"log"
	
	"github.com/katatrina/SWD392_NET1701_GroupIntern_BE/api"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern_BE/db/sqlc"
	"github.com/katatrina/SWD392_NET1701_GroupIntern_BE/internal/util"
	
	_ "github.com/lib/pq"
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
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	
	// Initialize database connection pool
	dbConn, err := sql.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	
	if err = dbConn.Ping(); err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	
	// Initialize database store
	store := db.NewStore(dbConn)
	
	// Initialize our HTTP server
	server := api.NewServer(store, config)
	
	// Start our HTTP server
	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
