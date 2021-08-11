package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) statusHandler(rw http.ResponseWriter, r *http.Request) {
	response := AppStatus{
		Status:      "Available",
		Environment: app.config.Env,
		Version:     version,
	}

	js, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(js)
}
