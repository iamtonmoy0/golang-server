package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("server started")
	// getReq()
	// postReq()
	getJson()
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
func postReq() {
	myurl := "https://google.com"
	// json
	requestBody := strings.NewReader(`
	{
		"name": "John Doe",
		"age": 35,
		"platform":"google.com"

	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// Get the body of the request
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func getJson() {
	url := "https://jsonplaceholder.typicode.com/posts"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyBytes))
}
