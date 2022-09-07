package jokes

type Repository interface {
	CreateNewJoke(newJoke *Joke) error
	GetJokes(offset, limit int) (*[]Joke, error)
	GetRandomJoke() (*Joke, error)
}
