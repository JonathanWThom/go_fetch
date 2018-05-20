// This program receives an endpoint as an argument, and returns the speed of its response,
// as JSON.
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	// needs to receive arg from api endpoint
	router := mux.NewRouter()
	router.HandleFunc("/speed/{url}", GetSpeed).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// private
func fetchUrl(url string) time.Duration {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		// handle it
	}
	// Use responses, like in this thread? https://stackoverflow.com/questions/30526946/time-http-response-in-go
	return time.Since(start)
}

func GetSpeed(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	time := fetchUrl(params["url"])
	json.NewEncoder(w).Encode(time)
}
