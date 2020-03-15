package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var b [5]string
	b[0] = "HELLO"
	b[1] = "WORLD"
	fmt.Println(b)
	fmt.Println(strings.Split("a,b,c", ","))
}
