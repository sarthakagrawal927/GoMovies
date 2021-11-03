package main

import (
	"errors"
	"net/http"
	"server/models"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print(errors.New("invalid movie id"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("get movie with id:", id)

	movie := models.Movie{
		ID:          id,
		Title:       "The Shawshank Redemption",
		Description: "Two imprisoned",
		Year:        1994,
		ReleaseDate: "10-10-1994",
		Runtime:     120,
		Rating:      "R",
		MPAAARating: "R",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
