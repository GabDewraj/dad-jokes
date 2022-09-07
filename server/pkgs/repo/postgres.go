package repo

import (
	"log"
	"math/rand"
	"time"

	"api-keys/pkgs/jokes"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewJokeRepo(dbConn *gorm.DB) jokes.Repository {
	return &repo{dbConn}
}

// CreateNewJoke implements jokes.Repository
func (r *repo) CreateNewJoke(newJoke *jokes.Joke) error {
	return r.db.Create(newJoke).Error
}

// GetJokes implements jokes.Repository
func (r *repo) GetJokes(offset int, limit int) (*[]jokes.Joke, error) {
	log.Print(offset, limit)
	var jokes []jokes.Joke
	if err := r.db.Scopes(paginate(offset,
		limit)).Find(&jokes).Error; err != nil {
		return nil, err
	}
	log.Print(jokes)
	return &jokes, nil
}

// GetRandomJoke implements jokes.Repository
func (r *repo) GetRandomJoke() (*jokes.Joke, error) {
	// Select all joke id's
	type Result struct {
		ID int
	}
	var result []Result
	if err := r.db.Table("jokes").Select("id").Scan(&result).Error; err != nil {
		return nil, err
	}
	// Generate a random int
	// Initialize the default Source to a deterministic state each iteration.
	// This allows each character to be different when selected.
	rand.Seed(time.Now().UnixNano())
	randomJokeId := rand.Intn(len(result) + 1)
	var jokeToReturn jokes.Joke
	jokeToReturn.ID = randomJokeId
	if err := r.db.First(&jokeToReturn).Error; err != nil {
		return nil, err
	}
	return &jokeToReturn, nil
}

// Custom func for pagination in GORM
// GORM uses the database/sql‘s argument placeholders to construct the SQL statement,
// which will automatically escape arguments to avoid SQL injection

func paginate(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}
