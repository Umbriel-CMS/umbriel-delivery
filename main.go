package main

import (
	"fmt"
	"net/http"
	"time"
)

var cache *Cache

func main() {
	cache = NewCache(5 * time.Minute)
	mux := http.NewServeMux()

	mux.HandleFunc("/receive", receiveDataHandler)
	corsWrappedMux := corsMiddleware(mux)

	http.HandleFunc("/receive", receiveDataHandler)

	fmt.Println("Umbriel Parse-Delivery is up on port 8080...")
	if err := http.ListenAndServe(":8080", corsWrappedMux); err != nil {
		panic(err)
	}
}
