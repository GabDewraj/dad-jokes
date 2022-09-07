package jokes

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
	return s.repo.GetJokes(offset, limit)
}

func (s *service) GetRandomJoke() (*Joke, error) {
	return s.repo.GetRandomJoke()
}
