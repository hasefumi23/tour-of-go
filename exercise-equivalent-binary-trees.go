package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Left -> Value -> Right の順で出力する
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		v1, v2 := <-ch1, <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	t1 := tree.New(1)
	t2 := tree.New(2)
	if Same(t1, t2) {
		fmt.Println("Same")
	} else {
		fmt.Println("Not Same")
	}
}
