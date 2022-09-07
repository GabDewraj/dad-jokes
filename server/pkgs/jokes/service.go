package jokes

import "log"

type Service interface {
	CreateNewJoke(newJoke *Joke) error
	GetJokes(offset, limit int) (*[]Joke, error)
	GetRandomJoke() (*Joke, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateNewJoke(newJoke *Joke) error {
	return s.repo.CreateNewJoke(newJoke)
}

func (s *service) GetJokes(offset, limit int) (*[]Joke, error) {
	jokes, err := s.repo.GetJokes(offset, limit)
	log.Print(jokes)
	return jokes, err
	// print("hello")
	// return nil, nil
}

func (s *service) GetRandomJoke() (*Joke, error) {
	return s.repo.GetRandomJoke()
}
