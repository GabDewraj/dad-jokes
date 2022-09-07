package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GabDewraj/dad-jokes/pkgs/jokes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockJokesService struct {
	mock.Mock
}

func (m *mockJokesService) CreateNewJoke(newJoke *jokes.Joke) error {
	args := m.Called(newJoke)
	return args.Error(0)
}

func (m *mockJokesService) GetJokes(offset, limit int) (*[]jokes.Joke, error) {
	args := m.Called(offset, limit)
	result := args.Get(0)
	return result.(*[]jokes.Joke), args.Error(1)
}

func (m *mockJokesService) GetRandomJoke() (*jokes.Joke, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*jokes.Joke), args.Error(1)
}

func TestCreateNewJoke200(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	reqbody := &map[string]interface{}{
		"body":   "dojcnwodcwnod",
		"author": "djncwdnc",
	}
	body, err := json.Marshal(reqbody)
	if err != nil {
		t.Fatal("Failed to marshal body")
		return
	}
	req := httptest.NewRequest(http.MethodPost, "/new/jokes", bytes.NewReader(body))
	res := httptest.NewRecorder()
	service := new(mockJokesService)
	newJoke := &jokes.Joke{
		Body:   "dojcnwodcwnod",
		Author: "djncwdnc",
	}
	service.On("CreateNewJoke", newJoke).Return(nil)
	handler := NewHandler(service)
	handler.CreateNewJoke(res, req)
	assertWithTest.Equal(http.StatusOK, res.Code)
}

func TestCreateNewJoke500(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	reqbody := &map[string]interface{}{
		"body":   "dojcnwodcwnod",
		"author": "djncwdnc",
	}
	body, err := json.Marshal(reqbody)
	if err != nil {
		t.Fatal("Failed to marshal body")
		return
	}
	req := httptest.NewRequest(http.MethodPost, "/new/jokes", bytes.NewReader(body))
	res := httptest.NewRecorder()
	service := new(mockJokesService)
	newJoke := &jokes.Joke{
		Body:   "dojcnwodcwnod",
		Author: "djncwdnc",
	}
	service.On("CreateNewJoke", newJoke).Return(errors.New("unknown"))
	handler := NewHandler(service)
	handler.CreateNewJoke(res, req)
	assertWithTest.Equal(http.StatusInternalServerError, res.Code)
}

func TestCreateNewJoke400(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	reqbody := &map[string]interface{}{
		"body":   323,
		"author": "djncwdnc",
	}
	body, err := json.Marshal(reqbody)
	if err != nil {
		t.Fatal("Failed to marshal body")
		return
	}
	req := httptest.NewRequest(http.MethodPost, "/new/jokes", bytes.NewReader(body))
	res := httptest.NewRecorder()
	service := new(mockJokesService)
	handler := NewHandler(service)
	handler.CreateNewJoke(res, req)
	assertWithTest.Equal(http.StatusBadRequest, res.Code)
}

func TestGenerateKey200(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	req := httptest.NewRequest(http.MethodPost, "/apikey", nil)
	res := httptest.NewRecorder()
	service := new(mockJokesService)
	handler := NewHandler(service)
	handler.GenerateApiKey(res, req)
	assertWithTest.Equal(http.StatusOK, res.Code)
}

func TestGetJokes200(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	req := httptest.NewRequest(http.MethodPost, "/new/jokes", nil)
	limit := 2
	offset := 3
	q := req.URL.Query()
	q.Add("limit", "2")
	q.Add("offset", "3")
	req.URL.RawQuery = q.Encode()
	res := httptest.NewRecorder()
	service := new(mockJokesService)
	jokes := &[]jokes.Joke{}
	service.On("GetJokes", offset, limit).Return(jokes, nil)
	handler := NewHandler(service)
	handler.GetJokes(res, req)
	assertWithTest.Equal(http.StatusOK, res.Code)
}

func TestGetRandomJoke200(t *testing.T) {
	t.Parallel()
	assertWithTest := assert.New(t)
	req := httptest.NewRequest(http.MethodPost, "/new/jokes", nil)
	res := httptest.NewRecorder()
	service := new(mockJokesService)
	joke := &jokes.Joke{}
	handler := NewHandler(service)
	service.On("GetRandomJoke").Return(joke, nil)
	handler.GetRandomJoke(res, req)
	assertWithTest.Equal(http.StatusOK, res.Code)
}
