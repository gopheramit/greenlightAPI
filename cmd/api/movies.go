package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gopheramit/greenlightAPI/internal/data"
	"github.com/gopheramit/greenlightAPI/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJson(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	movie := &data.Movie{
		Title:   input.Title,
		Runtime: input.Runtime,
		Genres:  input.Genres,
		Year:    input.Year,
	}
	v := validator.New()
	// v.Check(input.Title != "", "title", "title must be provided")
	// v.Check(len(input.Title) < 500, "title", "title must not be 500 bytes long")
	// v.Check(input.Year != 0, "year", "year must be provided")
	// v.Check(input.Year > 1888, "year", "year must be greater that 1888")
	// v.Check(input.Year < int32(time.Now().Year()), "year", "year must not be in future")
	// v.Check(input.Runtime != 0, "runtime", "runtime must be provided")
	// v.Check(input.Runtime > 0, "runtime", "runtime must be a possitve interger")
	// v.Check(input.Genres != nil, "genres", "genres must be provided")
	// v.Check(len(input.Genres) >= 1, "genres", "at least one genre to be provided")
	// v.Check(len(input.Genres) <= 5, "genres", "must not contain more than 5 genre")
	// v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate value")
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanc",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}
	err = app.writeJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		// app.logger.Println(err)
		// http.Error(w, "The server encountered problem and could not complete the request", http.StatusInternalServerError)
		app.serverErrorResponse(w, r, err)
	}

	// fmt.Fprintf(w, "show details of  movie id %d", id)
}
