package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"
)

func fetchPageBlocks() ([]InputBlockData, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	sponsorApiAddress := os.Getenv("API_UMBRIEL_ADDRESS_PORT")

	url := sponsorApiAddress + "/page-blocks"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error when making a request for %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the API responded with the status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var pageBlocks []InputBlockData
	err = json.Unmarshal(bodyBytes, &pageBlocks)
	if err != nil {
		return nil, fmt.Errorf("error when parsing JSON: %w", err)
	}

	return pageBlocks, nil
}
