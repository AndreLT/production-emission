package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"production-emission/datastore"
	"runtime"
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

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	port := os.Getenv("PORT")
	port = ":" + string(port)
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := mux.NewRouter()
	log.Println("production and emission api")
	PrintMemUsage()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/country/{iso}", SearchByIso).Methods(http.MethodGet)
	api.HandleFunc("/collection", SearchCollection).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe(port, r))
}
