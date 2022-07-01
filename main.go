package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" //import gorilla/mux toolkit
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	//set header
	w.Header().Set("Content-Type", "application/json")
	//encode response ie w to json
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	//seting header
	w.Header().Set("Content-Type", "application/json")

	//getting the params
	params := mux.Vars(r)

	//TODO: validate the param

	//loop through the movies slice
	for _, item := range movies {
		if item.Id == params["id"] {
			//encoding our item as json
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, req *http.Request) {
	//seting header
	w.Header().Set("Content-Type", "application/json")

	//TODO: validate the request body

	//movie instance
	var newMovie Movie

	//decode the request body to a reference of newmovie
	_ = json.NewDecoder(req.Body).Decode(&newMovie)

	//set id of newmovie to a random number that is converted to string
	newMovie.Id = strconv.Itoa(rand.Intn(100000000000))

	//append newMovie to movie
	movies = append(movies, newMovie)

	//encode movies to json
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	//set header
	w.Header().Set("Content-Type", "application/json")
	//get params
	params := mux.Vars(r)

	//TODO: validate the param

	//loop through the movies slice
	for i, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

	//encoding movie as json
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set content type
	w.Header().Set("Content-Type", "application/json")

	//get params
	params := mux.Vars(r)

	//TODO: validate the param

	for index, item := range movies {
		if item.Id == params["id"] {
			//delete movie with id
			movies = append(movies[:index], movies[index:]...)

			//movie instance
			var updatedMovie Movie
			//decode the request body to a reference of newmovie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			//set id of newmovie to a random number that is converted to string
			updatedMovie.Id = params["id"]
			//append newMovie to movie
			movies = append(movies, updatedMovie)
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}

	}

}

//a slice of Movie
var movies []Movie

func main() {
	movies = append(movies, Movie{Id: "1", Isbn: "552512",
		Title:    "young Guns",
		Director: &Director{Firstname: "Mr andrews", Lastname: "chan"}})

	movies = append(movies, Movie{Id: "2", Isbn: "452512",
		Title:    "peace on mars",
		Director: &Director{Firstname: "Mr James", Lastname: "orbit"}})

	//initial a new router instance
	route := mux.NewRouter()

	route.HandleFunc("/movies", getMovies).Methods("GET")
	route.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	route.HandleFunc("/movies", createMovie).Methods("POST")
	route.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	route.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//print to console
	fmt.Printf("starting server at port 8000\n")

	//starting the server
	log.Fatal(http.ListenAndServe(":8000", route))
}
