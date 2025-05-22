package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/foo", nil)
	rr := httptest.NewRecorder()

	apiHandler(rr, req)

	if got, want := rr.Code, http.StatusOK; got != want {
		t.Errorf("status: got %d, want %d", got, want)
	}
	if got, want := rr.Body.String(), "Hello from API"; got != want {
		t.Errorf("body: got %q, want %q", got, want)
	}
}
