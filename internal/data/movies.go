package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title,omitempty"`
	Year      int32     `json:",omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string
	Version   int32
}
