package main

import (
	"log"
	"net/http"
	"os"

	"api-keys/cmd/config"
	"api-keys/pkgs/api"
	"api-keys/pkgs/jokes"
	"api-keys/pkgs/repo"

	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
)

var env = os.Getenv("ENV")

func main() {
	newConfig, err := config.NewConfig(env)
	if err != nil {
		panic(err)
	}
	log.Print("Header", newConfig.ApiKeyConfig.Header)
	dbConn, err := config.NewDB(&newConfig.DbConfig, 3)
	if err != nil {
		panic(err)
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
