package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) getOneMovie(rw http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println("Invalid id parameter")
		app.errorJSON(rw, err)
		return
	}

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(rw, http.StatusOK, movie)
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *Application) getAllMovies(rw http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(rw, err)
	}

	err = app.writeJSON(rw, http.StatusOK, movies)
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *Application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GenresAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
