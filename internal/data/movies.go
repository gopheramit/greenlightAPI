package data

import (
	"time"

	"github.com/gopheramit/greenlightAPI/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title,omitempty"`
	Year      int32     `json:",omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string
	Version   int32
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "title must be provided")
	v.Check(len(movie.Title) < 500, "title", "title must not be 500 bytes long")
	v.Check(movie.Year != 0, "year", "year must be provided")
	v.Check(movie.Year > 1888, "year", "year must be greater that 1888")
	v.Check(movie.Year < int32(time.Now().Year()), "year", "year must not be in future")
	v.Check(movie.Runtime != 0, "runtime", "runtime must be provided")
	v.Check(movie.Runtime > 0, "runtime", "runtime must be a possitve interger")
	v.Check(movie.Genres != nil, "genres", "genres must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "at least one genre to be provided")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genre")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate value")
}
