package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var movies []Movie

func main() {
	initMovie()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /movies", getMovies)
	mux.HandleFunc("GET /movie/{id}", getMovieByid)
	mux.HandleFunc("POST /movie", createMovie)
	mux.HandleFunc("PUT /movie/{id}", updateMovie)
	mux.HandleFunc("DELETE /movie/{id}", deleteMovie)
	err := http.ListenAndServe("localhost:8080", mux)
	log.Println("server is running on 8080 port")
	if err != nil {
		log.Fatal("Fail to start server")
	}
}
func initMovie() {

	movies = append(movies, Movie{
		ID:    1,
		Isbn:  "ISBN1",
		Title: "Iron man",
		Director: &Director{
			Firstname: "Tony",
			Lastname:  "Stark",
		},
	})
	movies = append(movies, Movie{
		ID:    2,
		Isbn:  "ISBN2",
		Title: "Maine pyar kiya",
		Director: &Director{
			Firstname: "Salman",
			Lastname:  "khan",
		},
	})
	movies = append(movies, Movie{
		ID:    3,
		Isbn:  "ISBN3",
		Title: "Maine pyar kiya -2",
		Director: &Director{
			Firstname: "Salman",
			Lastname:  "khan",
		},
	})
}
func getJson(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("failed to encode the data :", err.Error())
	}
	return string(jsonData)
}
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, getJson(movies))

}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Movie not found", http.StatusBadRequest)
		return
	}

	_, id, err = getMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movies = append(movies[:id], movies[id+1:]...)
	fmt.Fprint(w, "movie deleted successfull")
}
func getMovie(id int) (*Movie, int, error) {
	for i, v := range movies {
		if v.ID == id {
			return &v, i, nil
		}
	}
	return nil, 0, errors.New("Movie not found")
}
func getMovieByid(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Fprint(w, getJson(err.Error()))
		return
	}
	movie, id, err := getMovie(id)
	if err != nil {
		log.Println("%v not found ", id)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, getJson(movie))
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}
	movie.ID = len(movies) + 1
	movies = append(movies, movie)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, getJson(movie))
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}
	var movie Movie
	err = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {

		http.Error(w, "invalid id Data", http.StatusBadRequest)
		return
	}
	found := false
	w.Header().Set("Content-Type", "application/json")
	for i, v := range movies {
		if v.ID == id {
			found = true
			movie.ID = id
			movies[i] = movie
			break

		}
	}
	if found {
		fmt.Fprint(w, getJson(movie))
	} else {
		fmt.Fprint(w, "Movie not found ")
	}

}

type Movie struct {
	ID       int       `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
