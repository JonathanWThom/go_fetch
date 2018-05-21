// This program receives an endpoint as an argument, and returns the speed of its response,
// as JSON.
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/speed/{url}", GetSpeed).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func fetchUrl(url string) float64 {
	fullUrl := "http://" + url
	start := time.Now()
	res, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	return time.Since(start).Seconds()
}

func GetSpeed(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var total float64
	for i := 0; i < 10; i++ {
		total += fetchUrl(params["url"])
	}
	avg := total / 10
	// This should return an json object of all tested values
	// Do them all at once (with a channel) rather than in sequence
	// Otherwise, it's a long API call
	json.NewEncoder(w).Encode(avg)
}
