package jokes

import (
	"time"
)

// Domain object
type Joke struct {
	ID        int       `json:"id,omitempty"`
	Body      string    `json:"body,omitempty"`
	Author    string    `json:"author,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
