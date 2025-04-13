package main

import (
	"fmt"
	"sync"
)

func main(){
	var count int
	increment := func(){
		count ++
	}

	var once sync.Once // *

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i<100 ; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)  // * that utilizes some sync primitives  internally to ensure that pnly one call to "DO" ever calls the function passed in --- even on the different goroutine
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

