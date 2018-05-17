// This program receives an endpoint as an argument, and returns the speed of its response,
// as JSON.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// needs to receive arg from api endpoint
	var time = fetchUrl("www.example.com")
	fmt.Println(time)
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
