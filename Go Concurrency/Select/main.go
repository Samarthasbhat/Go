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