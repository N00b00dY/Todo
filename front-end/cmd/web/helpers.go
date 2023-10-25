package main

import (
	"encoding/json"
	"net/http"
)

// writeJSON takes a response status code an arbitrary data nand writes a json response to the client
func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "apllication/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// throwJSONError takes an error and optionally a response status code and generates and
// sends a json error response
func (app *Config) throwJSONError(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	// create an json object
	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)

}
