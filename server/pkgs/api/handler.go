package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"api-keys/pkgs/jokes"
	"api-keys/pkgs/secrets"
)

type handler struct {
	service jokes.Service
}

type Handler interface {
	GenerateApiKey(res http.ResponseWriter, req *http.Request)
	GetRandomJoke(res http.ResponseWriter, req *http.Request)
	CreateNewJoke(res http.ResponseWriter, req *http.Request)
	GetJokes(res http.ResponseWriter, req *http.Request)
	HealthCheckHandler(res http.ResponseWriter, req *http.Request)
}

func NewHandler(service jokes.Service) Handler {
	return &handler{service: service}
}

func (h *handler) GetRandomJoke(res http.ResponseWriter, req *http.Request) {
	joke, err := h.service.GetRandomJoke()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(res).Encode(&joke); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetJokes(res http.ResponseWriter, req *http.Request) {
	limit, err := strconv.Atoi(req.URL.Query().Get("limit"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	offset, err := strconv.Atoi(req.URL.Query().Get("offset"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print(limit, offset)
	retrievedJokes, err := h.service.GetJokes(offset, limit)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(res).Encode(retrievedJokes); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) CreateNewJoke(res http.ResponseWriter, req *http.Request) {
	var newJokeReq jokes.Joke
	if err := json.NewDecoder(req.Body).Decode(&newJokeReq); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateNewJoke(&newJokeReq); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	respBody := fmt.Sprintf("Joke with id %d and author %s saved successfully",
		newJokeReq.ID, newJokeReq.Author)
	_ = json.NewEncoder(res).Encode(&respBody)
}

func (h *handler) GenerateApiKey(res http.ResponseWriter, req *http.Request) {
	key := secrets.AddSecret()
	message := fmt.Sprintf("Here is your key %s", key)
	_, _ = res.Write([]byte(message))
}

func (h *handler) HealthCheckHandler(res http.ResponseWriter, req *http.Request) {
	log.Print("Health check is successful...")
	res.WriteHeader(200)
	return
}
