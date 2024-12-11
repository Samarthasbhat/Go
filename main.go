package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr [5]int
	arr[0] = 10
	fmt.Println(arr[0])

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	Slices()
	CreateArray()
	UseMake()
	CapMake()
	LenCap()
	Array()
	slice := []int{4, 2, 9, 7, 1, 5, 6}
	max, min := findMaxMin(slice)
	fmt.Println()
	fmt.Println("Maximum:", max)
	fmt.Println("Minimum:", min)
	slices := []int{4, 2, 9, 7, 1, 5, 6}

	Sort(slices)
}

// Slices are reference type

func Slices() {
	var s []int
	s = append(s, 10, 20, 30)
	fmt.Println(s)
}

// Creating a slice from an array

func CreateArray() {
	arrC := [5]int{1, 2, 3, 4, 5}
	sC := arrC[1:4]
	fmt.Println(sC)
}

func UseMake() {
	sM := make([]int, 3, 4)
	fmt.Println(len(sM), cap(sM))
}

//Capacity in make

func CapMake() {
	c := make([]int, 3, 5)
	fmt.Println(cap(c))
}

//The length (len) must always be ≤ capacity (cap).

func LenCap() {
	s := make([]int, 2, 4)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// Append elements
	s = append(s, 1, 2)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// Exceeding Capacity
	s = append(s, 3)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

// Write a program to Reverse an Array

func Array() {
	var arr = [5]int{1, 2, 3, 4, 5}
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}
}

// Find a maximum and minimum element in Silces

func findMaxMin(slice []int) (int, int) {

	if len(slice) == 0 {
		panic("slice is empty")
	}

	max := slice[0]
	min := slice[0]

	for _, value := range slice {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return max, min
}

// import sort package
func Sort(slices []int) {
	//	sort.Ints(slices) // ascending order
	sort.Sort(sort.Reverse(sort.IntSlice(slices)))
	//	fmt.Println("Sorted slice:", slices)
	fmt.Println("Sorted slice in descending", slices)
}
