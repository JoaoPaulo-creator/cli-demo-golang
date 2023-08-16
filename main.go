package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"

	//"io"
	"net/http"

	"example.com/teste/structs"
)

var (
	help    bool
	request bool
	resp    bool
	list    bool
)

func main() {

	fmt.Println("Welcome to the Go CLI App!")

	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&request, "request", false, "Make an http request")
	flag.BoolVar(&resp, "response", false, "Make an http request to list something")
	flag.BoolVar(&list, "list", false, "Make an http request to list something")
	flag.Parse()

	endpoint := "http://localhost:3001/posts"

	bodyRequest := structs.Post{
		Title:   "Golang cli",
		Content: "post criado pela cli feita em golang",
	}

	jsonData, err := json.Marshal(bodyRequest)
	if err != nil {
		panic(err)
	}

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if request {
		responseBody := string(bodyBytes)
		fmt.Println(responseBody)
	}

	if resp {
		// Request a list of posts
		resp, err := http.Get(endpoint)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var posts []structs.ResponsePost
		err = json.NewDecoder(resp.Body).Decode(&posts)
		if err != nil {
			panic(err)
		}

		fmt.Println(posts)
	}

	if list {

		link := "http://localhost:3001/posts/47927970-cf5c-4bd9-8738-9100364051d1"
		// Request a list of posts
		resp, err := http.Get(link)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var posts structs.ResponsePostId
		err = json.NewDecoder(resp.Body).Decode(&posts)
		if err != nil {
			panic(err)
		}

		fmt.Println(posts.Title)
	}

	if flag.NArg() > 0 {
		fmt.Printf("Arguments: %s\n", flag.Args())
	}
}
