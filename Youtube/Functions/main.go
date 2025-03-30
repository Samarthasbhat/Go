//  Learn about function and defer key word


	// We can declare functions within a functions
	// Functions are always structural typing
	//  Formal parameters (a&b): func do(a, b int) int {...}

	// There is no "Pass by reference in GO"
	// By value parameters : num, bool, arr, struct
	// By reference parameters: pointer(&x), strings, slices, maps, channels
	// A function may call itself: the trick is knowing when to stop (RECURSION)
	// Recursion is solwer but it usefull in the Graphs, Tree structures 

	//The  defer statement caputres a fucntion call to run later 

	package main 

import (
	"fmt"
)

// Pass By Value
func do(m1 map[int]int) {
	m1[3] = 0
	m1 = make(map[int]int)
	m1[4]=4
	// fmt.Printf("b@ %p\n", b)
	fmt.Println("m1", m1)
}

func main(){

	m := map[int]int{4:1, 7:2, 8:3}
	// a := [3]int{1,2,3} 
	// v := do(m)

	do(m)
	fmt.Println("m", m)
}