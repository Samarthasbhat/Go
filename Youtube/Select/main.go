package main

import (
	// "fmt"
	"time"
	"log"
)


func main () {
	chans := []chan int{
		make(chan int),
		make(chan int),

	}

	for i := range chans{
		go func(i int, ch chan<- int){
			for{
				time.Sleep(time.Duration(i)* time.Second)
				ch <- i
		}
		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++{  // Poling 
		select{
		case m0 := <-chans[0]:
			log.Println("read from channel ", m0)
		case m1 := <-chans[1]:
			log.Println("read from channel ", m1)
		}
	}
}