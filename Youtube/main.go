package main

import (
	"fmt"
	"os"  // required for the stdin and stderr
	// "youtube/hello" // Importing the package correctly
)

func main() {


	// fmt.Println(hello.Say(os.Args[1:])) // Calling the Say function

	var sum float64
	var n int

	for {
		var val float64

		 _, err := fmt.Fscanln(os.Stdin, &val)
			if err != nil{
				break
		}
		sum  += val
		n++
	}

	if n == 0{
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}
	fmt.Println("The average is", sum/float64(n))
}
