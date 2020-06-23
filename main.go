package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	// add articles route and map it to
	// returnAllArticles function
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}
	handleRequests()
}

// Article struct
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles declare global Articles array
var Articles []Article
