package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "OK",
		Environment: app.config.env,
		Version:     version,
	}
	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.logger.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
