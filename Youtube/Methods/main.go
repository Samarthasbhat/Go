package main

// Value receiver: a copy of the value is passed to the method.
// Pointer receiver: a pointer to the value is passed to the method.
// Value receiver: the method can modify the value, but the original value is not modified.
// Pointer receiver: the method can modify the value, and the original value is modified.

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"math"
)

type Point struct{
	X, Y float64
}

type Line struct{
	Begin, End Point
}

type Path []Point

type Distancer interface{
	Distance() float64
}

func (l Line) Distance() float64{
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func (p Path) Distance() (sum float64){
	for i:= 1; i< len(p); i++{
		sum += Line{p[i-1], p[i]}.Distance() // Line literal we cannot take a distance of a line
	}
	return sum
}

func PrintDistance(d Distancer){
	fmt.Println(d.Distance())
}

func (l Line) ScaleBy(f float64) Line{ // Pointer receiver
	l.End.X += (f-1) * (l.End.X - l.Begin.X)
	l.End.Y += (f-1) * (l.End.Y - l.Begin.Y)

	return Line{l.Begin, Point{l.End.X, l.End.Y}} // return a new line with the scaled end point

}

type Counter int

// value receiver

func(c Counter) Increment(){
	c++
}

// Pointer receiver

func (C *Counter) IncrementP(){
	*C++ // dereference the pointer to increment the value

}

type IntSlice []int

func (is IntSlice) String() string{
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, ";") + "]"
}



type ByteCounter int

//Byte counter is Writer 

func(b *ByteCounter) Write(p []byte) (int, error){ // p is buffer 
	l := len(p)
	*b += ByteCounter(l) // increment the counter by the length of the buffer
	return l, nil // return the length of the buffer	
}

func main(){
	var v IntSlice = []int{1,2,3}

	var s fmt.Stringer = v

	for i,x := range v{
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n",v) // type of the interface
	fmt.Printf("%T %[1]v\n",s) // if it is a stringer it use string

// Value receiver vs Pointer receiver


var c ByteCounter
f1, _ := os.Open("test.txt")
f2 := &c

n, _ := io.Copy(f2, f1)

fmt.Println("Copied", n, "bytes")
fmt.Println(c)

var c1 Counter = 3
c1.Increment()
fmt.Println("Original value of c1:", c1) // 3, because c1 is a value receiver
fmt.Println(c1) // 3, because c1 is a value receiver

var c2 Counter = 10
c2.IncrementP()
fmt.Println("Original value of c2:", c2) // 11, because c2 is a pointer receiver
fmt.Println(c2)

// Line example
side := Line{Point{1,2}, Point{4,6}}
perimeter := Path{{1,1}, {5,1}, {5,4}, {1,1}}

side.ScaleBy(2.9) // side is a value receiver, so it does not change the original value

PrintDistance(side)
PrintDistance(perimeter) // Instance of Path implements the Distancer interface


fmt.Println(Line{Point{1,2}, Point{4,6}}.ScaleBy(2).Distance()) // Instance of Line implements the Distancer interface
}



