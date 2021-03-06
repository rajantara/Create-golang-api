package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`

}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request)  {

	articles := Articles {
		Article{Name:"rajan", Email:"rajan@gmail.com", Phone:"0823547556"},
	}

	fmt.Println("Endpoint hit  All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"home page endpoint hit")
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Test Endpoint Worked")
}

func handleRequest(){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
