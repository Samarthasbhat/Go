package main

// Race conditions

import (
	"fmt"
	//"time"
)


// func main(){
// 	var data int

// go func(){
// 	data++
// }()
// time.Sleep(1*time.Second) // This is bad!
// if data == 0 {
// 	fmt.Printf("The value is %v.\n", data)
// }
// }


// Memory Access Sync

func main(){
	var data int
	go func(){data++}()
	if data == 0{
		fmt.Println("The value is 0.")
	}else {
		fmt.Printf("The value is %v.\n", data)
	}
}

// Critical sections in this cases
// 1. goroutine incrementing data var.
// 2. if condition check whether the value of data is 0.
// 3.fmt.Printf which retrieves the value of data for output.