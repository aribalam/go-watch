package main

import (
	"backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) getOneMovie(rw http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println("Invalid id parameter")
		app.errorJSON(rw, r, err)
		return
	}

	movie := models.Movie{
		ID:          id,
		Title:       "Some Title",
		Description: "Some Description",
		Year:        2021,
		ReleaseDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
		Runtime:     100,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(rw, http.StatusOK, movie)
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *Application) getAllMovies(rw http.ResponseWriter, r *http.Request) {

}
