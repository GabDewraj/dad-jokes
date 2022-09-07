package main

import (
	"log"
	"os"

	"github.com/GabDewraj/dad-jokes/db/cmd/config"
)

var env = os.Getenv("APP_ENV")

func main() {
	log.Print("This is the app env var ", env)
	newConfig, err := config.NewConfig(env)
	if err != nil {
		panic(err)
	}

	// Connect to the database
	dbConn, err := config.NewDB(&newConfig.DbConfig, 2)
	if err != nil {
		panic(err)
	}
	// Drop any existing tables
	if err := config.DropTables(dbConn); err != nil {
		panic(err)
	}

	// Create db tables
	if err := config.CreateTables(dbConn); err != nil {
		panic(err)
	}
	log.Print("DB is initialised")
}
