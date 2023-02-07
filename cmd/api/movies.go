package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gopheramit/greenlightAPI/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorRespone(w, r, http.StatusBadRequest, err.Error())
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
