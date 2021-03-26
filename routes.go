package main

import (
	"encoding/json"
	"net/http"
	"production-emission/loader"
	"strings"

	"strconv"

	"github.com/gorilla/mux"
)

func SearchCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := make([]*loader.EmissionYear, 0, 0)

	collection, err := getCollectionParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter:}`))
		return
	}

	limit, err := getLimitParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter:}`))
		return
	}

	for _, v := range collection {
		current := *books.SearchByIso(v, limit)
		data = append(data, current...)
	}

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

func getCollectionParam(r *http.Request) ([]string, error) {
	collection := make([]string, 0, 0)
	queryParams := r.URL.Query()
	c := queryParams.Get("countries")
	if c != "" {
		split := strings.Split(c, "-")
		collection = split
	}
	return collection, nil
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
