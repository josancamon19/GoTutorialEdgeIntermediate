package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global articles array
// that we can then populate in our main function
// to simulate a database
var articles []Article

func main2() {
	articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles", returnAllArticles).Methods("GET")
	router.HandleFunc("/articles", createArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", returnArticleById).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	articleData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Article data not found")
	}
	var article Article
	err = json.Unmarshal(articleData, &article)
	if err != nil {
		fmt.Fprintf(w, "Please send a valid article data")
	}
	articles = append(articles, article)
	w.WriteHeader(http.StatusCreated)
}

func returnArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["id"]
	for _, article := range articles {
		if article.Id == articleId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(article)
		}
	}
	fmt.Println("Endpoint Hit: returnAllArticles")
	w.WriteHeader(http.StatusNotFound)
}
