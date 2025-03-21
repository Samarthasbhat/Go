package main

import "fmt"

func main(){
	s := "ẽlite"

	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s)) // UTF-8 extension
	fmt.Printf("%8T %[1]v\n", []byte(s)) // 3 bytes were used for 
	fmt.Printf("length of string %d\n", len(s))

	b := []byte(s)
	fmt.Printf("8T %[1]v %d\n", b, len(b))
}