package main

import (
	"encoding/json"
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

func SearchByIso(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if val, ok := pathParams["iso"]; ok {
		data := *books.SearchByIso(val, 10)
		y, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(y)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "not found"}`))
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
	log.Fatalln(http.ListenAndServe(":8080", r))
}
