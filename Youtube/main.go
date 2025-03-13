package main

import (
	"fmt"
	"os"
	"youtube/hello" // Importing the package correctly
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(hello.Say(os.Args[1])) // Calling the Say function
	} else {
		fmt.Println(hello.Say("World"))
	}
}
