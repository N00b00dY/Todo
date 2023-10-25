package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handlers(t *testing.T) {
	var Tests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/fisch", http.StatusNotFound},
	}

	var app Config

	app.dbServiceName = ""
	routes := app.routes()

	//create a test server
	ts := httptest.NewServer(routes)
	defer ts.Close()

	// range through test data
	for _, tt := range Tests {
		t.Run(tt.name, func(t *testing.T) {
			// create a new request
			resp, err := ts.Client().Get(ts.URL + tt.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			// check the status code
			if status := resp.StatusCode; status != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatusCode)
			}
		})
	}

}
