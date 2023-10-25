package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ThrowJSONError(t *testing.T) {

	recorder := httptest.NewRecorder()
	// create a new Config
	app := &Config{}

	// trigger the throwJSONError function
	errMessage := "Test error message"
	err := app.throwJSONError(recorder, errors.New(errMessage), http.StatusNotFound)

	// check the status code
	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status %d; got %d", http.StatusNotFound, recorder.Code)
	}

	// check the response body
	expected := `{"error":true,"message":"Test error message"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected body to be %s; got %s", expected, recorder.Body.String())
	}

	// check the error
	if err != nil {
		t.Errorf("Expected error to be nil; got %v", err)
	}
}
