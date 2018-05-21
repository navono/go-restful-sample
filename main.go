package main

import (
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/navono/go-RESTful-sample/RomanNumerals"
)

func main() {
	useNewServeMux()
}

func romanNumeralsServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(RomanNumerals.Numerals[number]))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
		}
	})

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func customServeMux() {
	mux := &CustomServeMux{}
	http.ListenAndServe(":8080", mux)
}

func useNewServeMux() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/int", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})

	newMux.HandleFunc("/float", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	http.ListenAndServe(":8080", newMux)
}
