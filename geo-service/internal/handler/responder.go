package handler

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	Success(w http.ResponseWriter, status int, data interface{})
	Error(w http.ResponseWriter, status int, err error)
}

type JSONResponder struct{}

func (j *JSONResponder) Success(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (j *JSONResponder) Error(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
