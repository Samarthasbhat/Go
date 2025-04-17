package main 

import (
	"fmt"
)

func main() {

	//  Empty vs Nil
	var s [] int

	t := []int{}  // Pointer struct

	u := make([]int, 5)  // Pointer to actual storage

	v := make([]int, 0, 5)   // If the slice should be 0 assign capacity.

	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil)


	//  Len vs Cap

	a := [3]int{1,2,3}
	b := a[:1]

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := b[0:2] // NONSENSE
	fmt.Println("c = ", c)

	fmt.Println(len(b))
	fmt.Println(cap(b))

	fmt.Println(len(c)) 
	fmt.Println(cap(c))  // Because the base is 'a'

}