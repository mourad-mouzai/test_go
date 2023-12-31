package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movie, err := app.models.DB.Get(id)
	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	// 1. Call the getAll method to retrieve all movies
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	
	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}	
}

// func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
	
// }

// func (app *application) InsertMovie(w http.ResponseWriter, r *http.Request) {
// }

// func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {
// }

// func (app *application) searchMovie(w http.ResponseWriter, r *http.Request) {
// }