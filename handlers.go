package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func receiveDataHandler(w http.ResponseWriter, r *http.Request) {
	pageBlocks, err := fetchPageBlocks()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when fetching page blocks: %v", err), http.StatusInternalServerError)
		return
	}

	// serializing
	jsonResponse, err := json.Marshal(pageBlocks)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error serializing page blocks to JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse) // send...
}
