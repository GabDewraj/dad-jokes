package main

import (
	"log"
	"net/http"
	"os"

	"api-keys/cmd/config"
	"api-keys/pkgs/api"
	"api-keys/pkgs/jokes"
	"api-keys/pkgs/repo"

	"github.com/gorilla/mux"
)

var env = os.Getenv("ENV")

func main() {
	newConfig, err := config.NewConfig(env)
	if err != nil {
		panic(err)
	}

	dbConn, err := config.NewDB(&newConfig.DbConfig, 3)
	if err != nil {
		panic(err)
	}
	repo := repo.NewJokeRepo(dbConn)
	service := jokes.NewService(repo)
	newHandler := api.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/jokes", newHandler.GetJokes).Methods(http.MethodGet)
	router.HandleFunc("/random/jokes", newHandler.GetRandomJoke).Methods(http.MethodGet)
	router.HandleFunc("/key", newHandler.GenerateApiKey)
	// Create a subrouter for a protected route
	protectedRoutes := router.PathPrefix("/new").Subrouter()
	protectedRoutes.Use(api.ApiKeyAuth())
	protectedRoutes.HandleFunc("/jokes", newHandler.CreateNewJoke).Methods(http.MethodPost)
	// Run the server with https protocal
	certPath := newConfig.Certificate.CertificateFilePath
	keyPath := newConfig.Certificate.KeyFilePath
	log.Print("The server is running on port 8080")
	log.Fatal(http.ListenAndServeTLS(":8080", certPath, keyPath, router))
}
