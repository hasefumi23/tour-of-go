package main

import "fmt"

func main() {
	i, j := 42, 270

	p := &i
	fmt.Println(p)
	*p = 21
	fmt.Println(&i)
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
