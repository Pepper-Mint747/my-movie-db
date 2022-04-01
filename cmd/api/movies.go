package main

import (
	"fmt"
	"github.com/Pepper-Mint747/my-movie-db/internal/data"
	"net/http"
	"time"
)

// createMovieHandler
// @desc    POST a single movie
//@route    POST /v1/movies
//@access   Public
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create  new movie")
}

// showMovieHandler
// @desc    GET a single movie
//@route    GET /v1/movies/:id
//@access   Public
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Create a new instance of the Movie struct, containing the ID extracted from
	// the URL
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   106,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	// Encode the struct to the JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
