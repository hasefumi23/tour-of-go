package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	newS := []int{1, 2, 3}
	s = append(s, newS...)
	printSlice(s)

	// the slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
