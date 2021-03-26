package main

import (
	"fmt"
	"log"
	"net/http"
	"production-emission/datastore"
	"time"

	"github.com/gorilla/mux"
)

var (
	books datastore.Emission
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	books = &datastore.Books{}
	books.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/country/{iso}", SearchByIso).Methods(http.MethodGet)
	api.HandleFunc("/collection", SearchCollection).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
