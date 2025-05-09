package main

import (
	"fmt"

	
)

// Ad hoc confinement ensures safe access to shared data by confining it to a single goroutine.
// Lexical confinement ensures data is confined to a specific lexical scope.

func main() {
    // Ad hoc confinement: the `data` slice is confined to the `loopData` goroutine.
    data := []int{0, 1, 2, 3} // Initialize the data slice with values

    loopData := func(handleData chan<- int) {
        defer close(handleData) // Close the channel when done

        for i := range data {
            handleData <- data[i] // Send data to the channel
        }
    }

    handleData := make(chan int) // Create a channel
    go loopData(handleData)      // Start the goroutine

    // Receive and print data from the channel
    for num := range handleData {
        fmt.Println(num)
    }

    // Lexical confinement: the `data1` slice is confined to the main function.
    data1 := []int{10, 20, 30, 40} // Initialize another data slice

    loopData1 := func(handleData1 chan<- int) {
        defer close(handleData1) // Close the channel when done

        for _, value := range data1 { // Value is used in the lexical scope of the loop
            handleData1 <- value // Send data to the channel
        }
    }

    handleData1 := make(chan int) // Create another channel
    go loopData1(handleData1)     // Start the goroutine

    // Receive and print data from the channel
    for num := range handleData1 {
        fmt.Println(num)
    }

    // for-select loop: This is a common pattern for handling multiple channels.


    // Sending iteration var out on the channel

    done := make(chan struct{}) // Create a done channel to signal completion
    stringStream := make(chan string) // Create a channel for strings
    for _, s := range []string{"a", "b", "c"} {
        select {
        case <- done:
            return 
            case stringStream <- s:
      }
    }

    // Looping infinitely until the done channel is closed
    for {
        select{
            case<-done:
                return
            default:
        }
    }
}