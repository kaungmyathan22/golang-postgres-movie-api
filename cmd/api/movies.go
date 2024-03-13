package main

import (
	"fmt"
	"net/http"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now, we simply
// return a plain-text placeholder response.
func (app *application) createMovieHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintln(w, "create a new movie")
	if err != nil {
		panic(err)
		return
	}
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve // the interpolated "id" parameter from the current URL and include it in a placeholder // response.
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	_, err = fmt.Fprintf(w, "show the details of movie %d\n", id)
	if err != nil {
		return
	}
}
