package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// models
type Course struct {
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
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("POST")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// home func
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world</h1>"))

}

// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get courses by id
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// Handle the error (e.g., return a bad request response)
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}
	for _, course := range courses {
		if course.ID == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found ")
	return
}
// create course