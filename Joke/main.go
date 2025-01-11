package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Joke structure based on JokeAPI's response
type Joke struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`     // For two-part jokes
	Delivery string `json:"delivery"` // For two-part jokes
	Joke     string `json:"joke"`     // For single-line jokes
	Error    bool   `json:"error"`    // Indicates if there was an issue
}

func main() {
	url := "https://v2.jokeapi.dev/joke/Any?blacklistFlags=nsfw,racist,sexist,explicit"

	// Make an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching joke:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var joke Joke
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Handle errors from the API
	if joke.Error {
		fmt.Println("API returned an error while fetching the joke.")
		return
	}

	// Print the joke based on its type
	if joke.Type == "single" {
		fmt.Printf("Here's a joke:\n%s\n", joke.Joke)
	} else if joke.Type == "twopart" {
		fmt.Printf("Here's a joke:\n%s\n%s\n", joke.Setup, joke.Delivery)
	} else {
		fmt.Println("Unexpected joke type!")
	}
}
