package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReverseProxy_Proxying(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(418)
		io.WriteString(w, "I am")
	}))
	defer upstream.Close()

	u := strings.TrimPrefix(upstream.URL, "http://")
	parts := strings.Split(u, ":")
	rp := NewReverseProxy(parts[0], parts[1])

	handler := rp.ReverseProxy(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("apiHandler should not be called for non-/api paths")
	}))

	req := httptest.NewRequest("GET", "/foo/bar", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if got, want := rr.Code, 418; got != want {
		t.Errorf("status: got %d, want %d", got, want)
	}
	if got, want := rr.Body.String(), "I am"; got != want {
		t.Errorf("body: got %q, want %q", got, want)
	}
}
