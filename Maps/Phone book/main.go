package main

import (
	"fmt"
)

func main() {
	Phone()
}

func Phone() {

	numbers := []string{"7283722394", "9302933290"}
	phoneBook := map[string]string{
		"John":  "7283722394",
		"Alice": "9302933290",
	}

	fmt.Println("Before deletion: ", phoneBook)

	//delete(numbers, "7283722394") // we cannot delete a slice using "delete"
	delete(phoneBook, "John")

	fmt.Println("After deletion: ", phoneBook)

	numbers = append(numbers, "1234567890")
	fmt.Println("Updated slice:", numbers)

	// add to map

	phoneBook["Bob"] = numbers[2]
	fmt.Println("Updated map:", phoneBook)

	// search number using name

	//searchName := "John"
	searchName := "Alice"

	if number, found := phoneBook[searchName]; found {
		fmt.Printf("Phone number for %s: %s\n", searchName, number)
	} else {
		fmt.Printf("No entry found for %s.\n", searchName)
	}

}
