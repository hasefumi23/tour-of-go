package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var n, twoBefore, oneBefore int

	return func() int {
		var ret int
		if n == 0 {
			ret = 0
		} else if n == 1 {
			ret = 1
			oneBefore = 1
		} else {
			ret = oneBefore + twoBefore
			twoBefore = oneBefore
			oneBefore = ret
		}

		n++
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
