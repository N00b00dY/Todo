package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NoSurve(t *testing.T) {
	// Create a sample "next" handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Call the NoSurve function with the sample "next" handler
	csrfWrappedHandler := NoSurve(nextHandler)

	// Create a test request
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Serve the request using the csrfWrappedHandler
	csrfWrappedHandler.ServeHTTP(w, req)

	// Check if the CSRF settings are applied as expected

	csrfCookie := w.Result().Cookies()[0]

	// validate the CSRF settings
	if csrfCookie.HttpOnly != true {
		t.Error("HttpOnly attribute not set as expected")
	}
	if csrfCookie.Path != "/" {
		t.Error("Cookie path not set as expected")
	}
	if csrfCookie.Secure != false {
		t.Error("Secure attribute not set as expected")
	}
	if csrfCookie.SameSite != http.SameSiteLaxMode {
		t.Error("SameSite attribute not set as expected")
	}

	// Check if the response status code is as expected
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d; got %d", http.StatusOK, w.Code)
	}
}
