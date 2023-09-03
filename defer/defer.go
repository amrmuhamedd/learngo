package main

import "fmt"

func add(a, b int) {
	sum := a + b
	fmt.Println("Sum: ", sum)
}

func message(msg string) {
	fmt.Println(msg)
}

func main() {

	message("Begin")

	defer message("End")
	add(10, 20)
	add(20, 30)
	add(30, 40)
	add(30, 40)
	add(30, 40)
	add(30, 40)
	add(30, 40)
	add(30, 40)
	add(30, 40)
}
