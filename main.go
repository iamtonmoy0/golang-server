package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	// routers
	r := mux.NewRouter()
	courses = append(courses, Course{
		ID:     1,
		Name:   "Go Programming",
		Price:  250,
		Author: &Author{"John Doe", "https://johndoe.com"},
	})
	courses = append(courses, Course{
		ID:     2,
		Name:   "Python Programming",
		Price:  300,
		Author: &Author{"Jane Smith", "https://janesmith.com"},
	})
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("POST")
	r.HandleFunc("/course-create", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course/update/{id}", updateOneCourse).Methods("PATCH")
	r.HandleFunc("/course/delete/{id}", deleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// middleware
func (c *Course) IsEmpty() bool {
	return c.Name == ""
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

}

// create course
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create new course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is empty")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Body is empty")
	}
	rand.Seed(time.Now().UnixNano())
	course.ID = rand.Intn(100)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// update one course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting the is
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	// looping through courses
	for index, course := range courses {
		if course.ID == id {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.ID = id
			courses = append(courses, course)
			json.NewEncoder(w).Encode("course updated")
		}
	}
	json.NewEncoder(w).Encode("failed to update course")
}

// delete courses
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, course := range courses {
		if course.ID == id {
			courses = append(courses[:index], courses[:index+1]...)
			json.NewEncoder(w).Encode("Course Deleted")
		}
	}
	json.NewEncoder(w).Encode("Failed to Delete Course")

}
