package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isdn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:";lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(movies)
	}

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}

}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}
func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}

	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isdn: "6456", Title: "hell", Director: &Director{FirstName: "john", LastName: "wick"}})
	movies = append(movies, Movie{ID: "2", Isdn: "5233", Title: "heaven", Director: &Director{FirstName: "taylor", LastName: "swift"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("server starting at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
