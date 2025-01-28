package main

// Deadlocks: Occurs when a goroutine locks a mutex and fails to unlock it (e.g., due to missing Unlock() in some code paths).
// Performance Bottleneck: Overusing mutexes can reduce the performance of your program as it serializes goroutines.
// Double Locking: Locking the same mutex twice in the same goroutine without unlocking causes a deadlock.

import (
	"fmt"
	"sync"
)

type Counter struct{
	mu sync.Mutex
	value int
}

func ( c *Counter) Increment(){
	c.mu.Lock() // Lock the mutex
	defer c.mu.Unlock() // Ensure unlock even if panic occurs
	c. value++
}

func (c *Counter) GetValue() int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main(){

	counter := Counter{}
	var wg sync.WaitGroup

	for i:=0; i<5; i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			fmt.Println(i)
			counter.Increment()
		}(i)
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter.GetValue())
}


