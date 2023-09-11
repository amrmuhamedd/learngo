package main

import "fmt"

func main() {
	var amrSlice = []int{1, 2, 3, 23}

	fmt.Println(amrSlice, len(amrSlice), cap(amrSlice))
}
