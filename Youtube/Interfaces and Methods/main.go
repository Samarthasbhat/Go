// An interface variable is nil until initialized
//  It really has two parts: 
// a value or pointer of some type
// a pointer to the type information so the correct actual method can be identified

package main

import (
	"fmt"
)

type errFoo struct {
	err error
	path string
}

func (e errFoo) Error() string{
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

func XYZ(a int) *errFoo{
	return nil
}

	func main() {

		//  err := XYZ(1) // err would be *errFoo
		var err error = XYZ(1) // BAD: interface gets a nil concrete value


		if err != nil {
			fmt.Println("oops")
	}else {
		fmt.Println("OK!")
	}

	// Currying
	add5 := add(10)

	result := add5(5) // result is 15
	fmt.Println(result) // Output: 15
}

// Pointer vs value receiver

// Curried function

func add(a int) func (int) int{
	return func(b int) int{
		return a + b
	}
}

