package main

import (
	"fmt"
	"net/http"
)

// @desc createMovieHandler
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

	//Otherwise, interpolate the movie ID in a placeholder response
	fmt.Fprintf(w, "show the details of movie %d\n", id)

}
