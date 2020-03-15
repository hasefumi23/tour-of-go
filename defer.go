package main

import "fmt"

func main() {
	if true {
		defer fmt.Println("world")
		fmt.Println("hey")
	}
	fmt.Println("hello")
}
