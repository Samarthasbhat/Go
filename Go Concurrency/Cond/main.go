package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	c := sync.NewCond(&sync.Mutex{})	// create condition using a standard sync.Mutex as the Locker
	queue := make([]interface{}, 0, 10) // slice which contain 10 elements


	removeFromQueue := func(delay time.Duration){
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}


	for i:=0; i< 10; i++{
		c.L.Lock() // Critical section while calling Locker
		for len(queue) == 2{
			c.Wait()
		}

		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second)
		c.L.Unlock()

	}

}