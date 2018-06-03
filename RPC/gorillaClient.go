package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Args holds arguments passed to JSON RPC service
type Args struct {
	ID string
}

// Book struct holds Book JSON structure
type Book struct {
	ID     string `"json:string, omitempty"`
	Name   string `"json:name, omitempty"`
	Author string `"json:author, omitempty"`
}

// JSONServer RPC Serviec
type JSONServer struct{}

// GiveBookDetail ...
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	// Read JSON file and load data
	raw, readerr := ioutil.ReadFile("./books.json")
	if readerr != nil {
		log.Println("error: ", readerr)
		os.Exit(1)
	}
	// Unmarshal JSON raw data into books array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error: ", marshalerr)
		os.Exit(1)
	}

	// Iterate over each book to find the given book
	for _, book := range books {
		if book.ID == args.ID {
			*reply = book
			break
		}
	}

	return nil
}

func main() {
	// Create a new RPC server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")

	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
