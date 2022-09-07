package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/GabDewraj/dad-jokes/cmd/config"
	"github.com/GabDewraj/dad-jokes/pkgs/api"
	"github.com/GabDewraj/dad-jokes/pkgs/jokes"
	"github.com/GabDewraj/dad-jokes/pkgs/repo"
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

var (
	// Variables can injected for aws RDS at image build time with secrets manager
	env        = os.Getenv("ENV")
	dbHost     = os.Getenv("DB_HOST")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
	dbConn     *gorm.DB
)

func main() {
	newConfig, err := config.NewConfig(env)
	if err != nil {
		panic(err)
	}
	// Establish Database connection in aws or local.
	if env == "aws" {
		portNum, _ := strconv.Atoi(dbPort)
		awsDBConfig := &config.DbConfig{
			Host:     dbHost,
			User:     dbUser,
			Password: dbPassword,
			Database: dbName,
			Port:     portNum,
		}
		dbConn, err = config.NewDB(awsDBConfig, 2)
		if err != nil {
			panic(err)
		}

	} else {
		dbConn, err = config.NewDB(&newConfig.DbConfig, 2)
		if err != nil {
			panic(err)
		}
	}
	repo := repo.NewJokeRepo(dbConn)
	service := jokes.NewService(repo)
	newHandler := api.NewHandler(service)

	router := mux.NewRouter()
	// CORS middleware for frontend consumers
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	// Health check
	router.HandleFunc("/health-check", newHandler.HealthCheckHandler)
	router.HandleFunc("/jokes", newHandler.GetJokes).Methods(http.MethodGet)
	router.HandleFunc("/random/jokes", newHandler.GetRandomJoke).Methods(http.MethodGet)
	router.HandleFunc("/apikey", newHandler.GenerateApiKey)
	// Create a subrouter for a protected route
	protectedRoutes := router.PathPrefix("/new").Subrouter()
	protectedRoutes.Use(api.ApiKeyAuth(newConfig.ApiKeyConfig.Header))
	protectedRoutes.HandleFunc("/jokes", newHandler.CreateNewJoke).Methods(http.MethodPost)
	// Run the server with https protocal
	certPath := newConfig.Certificate.CertificateFilePath
	keyPath := newConfig.Certificate.KeyFilePath
	log.Print("The server is running on port 8080")
	log.Fatal(http.ListenAndServeTLS(":8080", certPath, keyPath, router))
}
