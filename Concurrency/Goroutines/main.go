package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello!")
}

func printMessage(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go sayHello() //Goroutine
	go printMessage("Goroutine 1")
	go printMessage("Goroutine 2")
	time.Sleep(3 * time.Second)

	fmt.Println("Main finished!")

}
