package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies []Movie

func moviesHandler(res http.ResponseWriter, req *http.Request) {
	method := req.Method

	if method == "GET" {
		b, err := json.Marshal(movies)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(res, "error: %s", err)
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.Write(b)
	} else if method == "POST" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(res, "error : %v", err)
			return
		}
		t := Movie{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(res, "error: %s", err)
			return
		}
		movies = append(movies, t)
		res.WriteHeader(http.StatusOK)
		b, err := json.Marshal(movies)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(res, "error: %s", err)
			return
		}
		res.Write(b)
		return
	}
}

func main() {
	http.HandleFunc("/movies", moviesHandler)
	err := http.ListenAndServe("localhost:8888", nil)
	log.Fatal(err)
}
