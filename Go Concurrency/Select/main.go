//  Select statement is the 'glue' that binds channels together.

package main

import (
	"fmt"
	"time"

)


func main() {
start := time.Now()

c:= make(chan interface{})

go func() {
	time.Sleep(5*time.Second) //After entering the select block 
	close(c)
}()

fmt.Println("Blocking on read ...")
select{
case <- c:
	fmt.Printf("Unblocked %v later.\n", time.Since(start))
}

c1Count, c2Count := RunSelectExample()
	fmt.Printf("c1Count: %d\n", c1Count)
	fmt.Printf("c2Count: %d\n", c2Count)


	Permanent()

	NotReady()

	Report()
}


//  Multiple channels being ready simultaneously seems interesting.

func RunSelectExample() (int, int){
	c1 := make(chan interface{}); close(c1)
	c2 := make(chan interface{}); close(c2)

	var c1Count, c2Count int

	for i:=1000;i >= 0;i--{
		select{   // Half-time 
		case <- c1:
			c1Count++
		case <- c2:
			c2Count++
		}
	}
	return c1Count, c2Count
}


//  What if there are never any channels that become ready? this is not permenant so we can use the time package

func Permanent() {
	var c<-chan int

	select{ 
	case <- c:  // This case statement will never become unblocked because we're reading from "nil" channel.
	case <-time.After(1*time.Second):
		fmt.Println("Timed out")
	}
}

//  When no channel is ready

func NotReady() {
	start := time.Now()
	var c1, c2 <-chan int 
	select{
	case <-c1:
	case<-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}

// Work while waiting for another goroutine to report a result

func Report(){
	done := make(chan interface{})

	go func ()  {
		time.Sleep(5*time.Second)
		close(done)
	}()
	workCounter  := 0
	loop:
	for {
		select{
		case <-done:
			break loop
		default:
		}

		// Simulate work
		workCounter++
		time.Sleep(1*time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop. \n", workCounter)
}