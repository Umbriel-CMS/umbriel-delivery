package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func receiveDataHandler(w http.ResponseWriter, r *http.Request) {
	inputData, err := fetchPageBlocks()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error when fetching page blocks: %v", err), http.StatusInternalServerError)
		return
	}

	outputData := make([]OutputBlockData, len(inputData))

	for i, inputBlock := range inputData {
		// Usar a função transformToOutputBlockDataWithCache para aproveitar o cache.
		outputData[i] = transformToOutputBlockDataWithCache(inputBlock, cache)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(outputData); err != nil {
		http.Error(w, fmt.Sprintf("Error serializing page blocks to JSON: %v", err), http.StatusInternalServerError)
	}
}
