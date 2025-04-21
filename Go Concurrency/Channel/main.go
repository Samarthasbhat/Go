package main

// Chan can be used for Unidirectional data - Sending or Recieving  (<-)

import (
	"fmt"
	"sync"
)

func main() {
	// var dataStream chan interface{} // Empty interface called as 'off type' interface
	// dataStream = make(chan interface{})

	// writeStream := make(chan<- interface{})
	// readStream := make(<-chan interface{})

	// <-writeStream
	// readStream <- struct{}{}

	// stringStream := make(chan string)
	// go func() {
	// 	stringStream <- "Hello channels"
	// }()
	// fmt.Println(<-stringStream)

	stringStream := make(chan string)
	go func() {
		// if 0 != 1 {
		// 	return
		// } 
		stringStream <- "hello channels"
		
	close(stringStream)
	}()

	// fmt.Println(<-stringStream)


for{
	salutation, ok := <- stringStream
	fmt.Printf("(%v): %v", ok, salutation)
	if !ok {
		fmt.Printf("(%v): %v\n", ok, salutation) // prints (false):
		break
	}

	fmt.Printf("(%v): %v\n", ok, salutation)
}
// 	You create a chan string

// You launch a goroutine

// The goroutine immediately returns due to if 0 != 1

// Nothing is ever sent to the channel

// Your main goroutine blocks on <-stringStream forever → deadlock



intStream := make(chan int)
go func() {
	defer close(intStream)
	for i:=1;i<=5;i++{
		intStream <-i
	}
}()

 for integer := range intStream {
	fmt.Printf("%v", integer)
 }


//  Unblocking multiple goroutines

begin := make(chan interface{})

var wg sync.WaitGroup

for i:= 0; i<5 ; i++{
	wg.Add(1)
	go func(i int){
		defer wg.Done()
		<- begin // Goroutines wait until it is told it can continue
		fmt.Printf("%v has begun\n", i)
	}(i)
} 
		fmt.Println("\n Unblocking goroutines....")
		close(begin)  // Unblocking all goroutines simultaneously 
		wg.Wait()
}
