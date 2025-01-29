package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

// Joke structure
type Joke struct {
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	file, err := os.Open("joke/joke.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read file content
	body, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Create a slice of jokes
	var jokes []Joke

	// Parse the JSON array into the slice
	err = json.Unmarshal(body, &jokes)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Check if there are any jokes
	if len(jokes) == 0 {
		log.Fatal("No jokes found in JSON file")
	}

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Select a random joke
	randomIndex := rand.Intn(len(jokes))
	joke := jokes[randomIndex]

	// Print the selected joke
	fmt.Println("Type:", joke.Type)
	fmt.Println("Setup:", joke.Setup)
	fmt.Println("Punchline:", joke.Punchline)
}
