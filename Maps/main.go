package main

import (
	"fmt"
	"strings"
	)


func main(){
	
	freq()
	makeMap()
}


// Creating a Map

// Using make:

func makeMap(){
	myMap := make(map[string]int)
	myMap["apple"] = 100
	myMap["Mango"] = 500

	fmt.Println("Initial map:", myMap)

	value, exists := myMap["apple"]

	if exists{
		fmt.Println("Value of apple:", value)
	}else{
		fmt.Println("Apple not found")
	}

	// update a value

	myMap["apple"] = 15

	// add new key
	myMap["Cherry"] = 20

	// delete a key
	delete(myMap,"Mango")


	// Iterate over the map
	fmt.Println("Updated map")
	for key, value := range myMap{
		fmt.Printf("Key : %s, Value: %d\n", key, value)
	}
}


// Create a map to store the frequency of words in a string.

func freq(){

	text := "this is a test this is only a test"

	//Split the string into words
	words := strings.Fields(text)

	// Create a map to store word frequencies
	count := make(map[string]int)


	for _, word := range words{
		count[word]++
	}

	fmt.Println(text)
	fmt.Println("Word Frequencies:")

	for word, freq := range count{
		fmt.Printf("%s : %d\n", word, freq)
	}
}