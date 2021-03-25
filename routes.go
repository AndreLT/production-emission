package main

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func SearchByIso(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	limit, err := getLimitParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter:}`))
		return
	}
	if val, ok := pathParams["iso"]; ok {
		data := *books.SearchByIso(val, limit)
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

func getLimitParam(r *http.Request) (int, error) {
	limit := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("limit")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return limit, err
		}
		limit = val
	}
	return limit, nil
}
