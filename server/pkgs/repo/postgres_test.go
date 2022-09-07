package repo

import (
	"testing"

	"github.com/GabDewraj/dad-jokes/cmd/config"
	"github.com/GabDewraj/dad-jokes/pkgs/jokes"
	"github.com/stretchr/testify/assert"
)

var dbConfig = config.DbConfig{
	Host:     "localhost",
	User:     "user",
	Password: "password",
	Database: "jokeDb",
	Port:     5432,
}

var dbConn, _ = config.NewDB(&dbConfig, 2)

func TestPaginationSuccess(t *testing.T) {
	assertWithTest := assert.New(t)
	var jokes []jokes.Joke
	offset := 3
	limit := 2
	dbConn.Scopes(paginate(offset, limit)).Select("author").Find(&jokes)
	assertWithTest.Nil(jokes, "paginated results of jokes")
}

func TestCreateNewJokeSuccess(t *testing.T) {
	assertWithTest := assert.New(t)
	repo := NewJokeRepo(dbConn)
	newJoke := &jokes.Joke{
		Body:   "why did the chicken cross the road ?, insert funny punchline #tired",
		Author: "John Doe",
	}
	err := repo.CreateNewJoke(newJoke)
	assertWithTest.Nil(err)
	assertWithTest.NotEqual(0, newJoke.ID)
}

func TestGetJokesSuccess(t *testing.T) {
	assertWithTest := assert.New(t)
	repo := NewJokeRepo(dbConn)
	limit := 2
	offset := 2
	jokes, err := repo.GetJokes(offset, limit)
	assertWithTest.Nil(err)
	assertWithTest.NotNil(jokes)
}

func TestGetRandomJokeSuccess(t *testing.T) {
	assertWithTest := assert.New(t)
	repo := NewJokeRepo(dbConn)
	joke, err := repo.GetRandomJoke()
	assertWithTest.Nil(err)
	assertWithTest.NotNil(joke)
}
