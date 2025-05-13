// An interface variable is nil until initialized
//  It really has two parts:
// a value or pointer of some type
// a pointer to the type information so the correct actual method can be identified

package main

import (
	"fmt"
	"math"
)

type errFoo struct {
	err  error
	path string
}

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

func XYZ(a int) *errFoo {
	return nil
}

func main() {

	//  err := XYZ(1) // err would be *errFoo
	var err error = XYZ(1) // BAD: interface gets a nil concrete value

	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("OK!")
	}

	// Currying
	add5 := add(10)

	result := add5(5)   // result is 15
	fmt.Println(result) // Output: 15

	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(p.Distance(q)) // Output: 5

	distanceFromP := p.Distance
	fmt.Printf("%T\n", distanceFromP)  // Output: func(*main.Point, main.Point) float64
	fmt.Printf("%T\n", Point.Distance) // Output: *func(*main.Point, main.Point) float64

	fmt.Println(distanceFromP(q))
}

// Pointer vs value receiver

// A method value with a value receiver copies the receiver

// Curried function

func add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// Interfaces in pratice

// Let consumers define interfaces
// Re-use standard interfaces wherever possible
// Keep interfaces declaration small
// Compose one-method interfaces into larger ones(if needed)
// Avoid coupling interfaces to particular types/implementations
// Accept interfaces, but return concrete types
// "Empty interfaces" are equal to null pointer in C++
