package main



import (
	"fmt"
)

// Define Interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Define Struct Rectangle
type Rectangle struct{
	Width, Height float64
}

// Implement Shape interface for Rectangle
func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64{
	return 2 *(r.Width + r.Height)
}

// Empty Interface
func describe(i interface{}){
	fmt.Printf("Value: %v, Type: %T\n", i,i)
}

// Type switch for assertion

func identifyType(i interface{}){
	switch v :=i.(type){
	case int:
		fmt.Println("Integer:",v)
	case string:
		fmt.Println("String:", v)
	case nil:
		fmt.Println(v)
	default:
		fmt.Println("Unknown Type")
	}
}


func main(){
	var s Shape
	s = Rectangle{4,5} 

	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())

	describe(42)
	describe("Hello, Go!")
	describe(3.14)
	describe(nil) // There is no NULL in GOLANG


	// Assert

	var i interface{} = 43

	num, ok:= i.(int)    // Assert if 'i' is of type int

	if ok {
		fmt.Println("Integer:",num)
	}else{
		fmt.Println("Not an integer")
	}

	//Identify
	identifyType(100)
	identifyType("GOLang")
	identifyType(nil)


	// Assertion failure

	var j interface{} = "Hello, Go!"

	num1 := j.(int)
	fmt.Println(num1)  //panic: interface conversion: interface {} is string, not int
}