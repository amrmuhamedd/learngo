package main

import "fmt"

func main() {
	var i interface{} = "jkj"

	// Retrieves the int value from the interface
	if x, ok := i.(int); ok {
		fmt.Println(x)
	} else {
		fmt.Println("Failed to retrieve an int value from the interface.")
	}

	// Tests if the interface has a string type and prints the value if found
	if y, ok := i.(string); ok {
		fmt.Printf("String value exists! The value is %v\n", y)
	} else {
		fmt.Println("No string value available!! Let's print zero value of the string i.e., empty string.")
	}

	// Tries to retrieve a float64 value but handles the error
	if z, ok := i.(float64); ok {
		fmt.Println(z)
	} else {
		fmt.Println("Failed to retrieve a float64 value from the interface.")
	}
}
