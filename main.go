package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	
)

func main() {
	fmt.Println("server started")
	getReq()
}

func getReq() {
	url := "https://google.com"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	// Get the body of the request
	fmt.Println("the status code is", response.StatusCode)
	fmt.Println("the length is", response.ContentLength)
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
