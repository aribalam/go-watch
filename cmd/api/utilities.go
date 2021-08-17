package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) writeJSON(rw http.ResponseWriter, status int, data interface{}) error {
	dat, err := json.Marshal(data)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(dat)

	return nil
}

func (app *Application) errorJSON(rw http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theErr := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(rw, http.StatusBadRequest, theErr)
}
