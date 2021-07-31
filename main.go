package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"` //unique id for every movie
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Movies Buzz Page")
}
func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for index, item := range movies {
		if item.Id == p["id"] {
			//updating movies slice by aappending the movies before id,after id
			//resulting in deletion of given id movie
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for _, item := range movies {
		if item.Id == p["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	//we will find the corresponding movie with same id, then delete it.
	//and add new movie with updated contents
	for index, item := range movies {
		if item.Id == p["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = p["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}
func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", homepage)
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Print("Starting Server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	movies = []Movie{
		Movie{Id: "1", Isbn: "101", Title: "Movies One", Director: &Director{Firstname: "Chris", Lastname: "Hiddleston"}},
		Movie{Id: "2", Isbn: "102", Title: "Movies Two", Director: &Director{Firstname: "Steve", Lastname: "Johnson"}},
		Movie{Id: "3", Isbn: "103", Title: "Movies Three", Director: &Director{Firstname: "Brave", Lastname: "Smith"}},
		Movie{Id: "4", Isbn: "104", Title: "Movies Four", Director: &Director{Firstname: "Jason", Lastname: "Roy"}},
		Movie{Id: "5", Isbn: "105", Title: "Movies Five", Director: &Director{Firstname: "David", Lastname: "Wisely"}},
	}
	handleRequest()
}
