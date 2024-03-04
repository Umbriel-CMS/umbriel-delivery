package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/receive", receiveDataHandler)

	fmt.Println("Umbriel Parse-Delivery is up on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
