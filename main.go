package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}

	// Default url parser lib
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("Url is in invalid format: %s\n", err)
		os.Exit(1)
	}

	// HTTP req
	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	err = json.Unmarshal(body, &words)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed\nPage: %s\nWords: %v", words.Page, strings.Join(words.Words, ", "))
}
