package main

import (
	"fmt"
)

// Ad hoc confinement is a powerful concurrency pattern in Go that ensures safe 
// access to shared data by confining it to a single goroutine. In this example, 
// the data slice is confined to the loopData goroutine, and channels are used 
// for safe communication between goroutines. This pattern is simple, efficient,
// and avoids the complexity of explicit synchronization.


func main() {
	data := make([]int, 4)


	loopData := func(handleData chan <- int){
		defer close(handleData)

		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}