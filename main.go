package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// models
type Courses struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Price  int     `json:"price"`
	Author *Author `json:"author"`
}
type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

// empty slice for db
var courses []Course

func main() {
	fmt.Println("server started")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/courses", getCourses)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world</h1>"))

}

// get all courses
func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get courses by id
func getCourse(w http.ResponseWriter, r *http.Request) {

}
