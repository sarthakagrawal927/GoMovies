package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data
	js, err := json.Marshal(wrapper)

	if err != nil {
		app.logger.Print(errors.New("invalid id"))
		app.errorJSON(w, err)
		return err
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error) {

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, http.StatusBadRequest, theError, "error")
}
