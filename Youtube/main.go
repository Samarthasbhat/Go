package main

import (
	"fmt"
	"os"
	"youtube/hello" // Importing the package correctly
)

func main() {

		fmt.Println(hello.Say(os.Args[1:])) // Calling the Say function

}
