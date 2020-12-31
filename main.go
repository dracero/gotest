package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "os"
    "github.com/gorilla/mux"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument	
   log.Fatal(http.ListenAndServe(GetPort(), router))
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
var port = os.Getenv("PORT")
// Set a default port if there is nothing in the environment
if port == "" {
		port = "4747"
 		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
 	}
 	return ":" + port
 }

func main() {
    Articles = []Article{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}
