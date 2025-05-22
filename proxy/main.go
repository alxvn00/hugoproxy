package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := chi.NewRouter()

	rp := NewReverseProxy("hugo", "1313")
	r.Use(rp.ReverseProxy)

	r.HandleFunc("/api", apiHandler)
	r.HandleFunc("/api/*", apiHandler)

	http.ListenAndServe(":8080", r)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

const content = ``

func WorkerTest() {
	t := time.NewTicker(1 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
