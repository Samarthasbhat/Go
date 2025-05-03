// It is a function controls the number of OS threads that will host so-called "work queues"

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

)

func main() {
	// Use all available CPU

	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup


	for i := 0; i<4; i++ {
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			fmt.Printf("Worker %d running on thread\n", id)
		}(i)
	}
	wg.Wait()


	// Homework

	totalFriends := 4

	for procs := 1; procs <= totalFriends; procs++{
		fmt.Println("\n Setting GOMAXPROCS to", procs)
		runtime.GOMAXPROCS(procs) // Set how many helpers (CPU Cores) can work at once

		var wg sync.WaitGroup
		start := time.Now()

		// start homework for all friends

		for i:= 1; i<= totalFriends; i++{
			wg.Add(1)
			go doHomework(i, &wg)
		}

		wg.Wait()
        duration := time.Since(start)
        fmt.Printf("Total time with GOMAXPROCS = %d: %v\n", procs, duration)

	}
}



func doHomework(friendID int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Friend %d started doing homework\n", friendID)
	time.Sleep(1 * time.Second) // Simulate time taken to  do homework
	fmt.Printf("Friend %d finished homework!\n", friendID)
}