//  Learn about function and defer key word


	// We can declare functions within a functions
	// Functions are always structural typing
	//  Formal parameters (a&b): func do(a, b int) int {...}

	// There is no "Pass by reference in GO"
	// By value parameters : num, bool, arr, struct
	// By reference parameters: pointer(&x), strings, slices, maps, channels
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