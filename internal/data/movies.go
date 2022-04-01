package data

import "time"

// Movie Anotate the Movie struct with struct tags to control how thw keys appear in
// the JSON encoded output
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"_"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
