package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("server started")
	r := mux.NewRouter()
	r.HandleFunc("/",home)
	
	log.Fatal(http.ListenAndServe(":8000",r))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world</h1>"))

}
