package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(GetPort(), nil))
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
    handleRequests()
}
