package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// articleHandler is a function handler
func articleHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars returns all path parameters as a map
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryParam["id"])
	fmt.Fprintf(w, "Got parameter category:%s!\n", queryParam["category"])
}

// CreateGorillaMuxHandler is a mux server
func CreateGorillaMuxHandler() {
	// Create a new router
	r := mux.NewRouter()

	// Query-based:
	r.HandleFunc("/articles", queryHandler)
	r.Queries("id", "category")
	// curl http://localhost:8000/articles/?id=123&category=books

	// Attach a path with handler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", articleHandler)

	// or we can:
	// r.Path("/articles/{category}/{id:[0-9]+}").HandlerFunc(articleHandler)

	// or add subpath:
	// s:=r.PathPrefix("/articles").Subrouter()
	// s.HandleFunc("{id}/settings", settingsHandler)
	// s.HandleFunc("{id}/details", detailsHandler)

	// Reverse mapping URL
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", articleHandler).Name("articleRoute")
	// url, _ := r.Get("articleRoute").URL("category", "books", "id", "123")
	// fmt.Printf(url.String())

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	// curl http://localhost:8000/articles/books/123
}
