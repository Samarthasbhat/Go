package main

import (
	"fmt"
	"sync"
)

//  Making a GUI how broadcast works

type Button struct {
	Clicked *sync.Cond
}

func subscribe(c *sync.Cond, fn func(), wg *sync.WaitGroup) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)

	go func() {
		goroutineRunning.Done()
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fn()
		wg.Done()
	}()
	goroutineRunning.Wait()
}

func main() {

	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	var clickRegistered sync.WaitGroup

	clickRegistered.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window")

	}, &clickRegistered)

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!!")

	}, &clickRegistered)

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse Clicked")

	}, &clickRegistered)

	button.Clicked.Broadcast()

	clickRegistered.Wait()
}
