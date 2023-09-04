package main

import "fmt"

func main() {
	// Declaring an array of integers with a length of 5
	var myArray [5]int

	// Assigning values to the elements of the array
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5

	// Accessing and printing array elements
	fmt.Println("Array:", myArray)
	fmt.Println("length:", len(myArray))
	fmt.Println("Element at index 2:", myArray[2])
}
