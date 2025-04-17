//  A closure is when a fucntion inside another function "closes over" one or more local variables of the outer functions

// The inner functions gets a reference to the outer functions variables
// Function returning a function
//  f := fib() 


package main


import (
	"fmt"
)

func fib() func() int{
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}


func do(d func()){
	d ()
}

func main() {
	f := fib()
	for x := f();  x<100; x = f() {
		fmt.Println(x)
	}

	var i int

	for i = 0;  i<4; i++ {
		v  := func (){
			fmt.Printf("%d @ %p\n", i, &i )
		}
		do(v)
	}

// Slice

	s :=  make([]func(), 4)

	for i :=0; i<4; i++{

		//  Fix the closures

		i  := i  // Closure capture i2 := i
			s[i] = func () {
				fmt.Printf("%d @ %p\n" , i ,&i)
			}
	}

	for i := 0; i<4; i++{
		s[i]()
	}

 }