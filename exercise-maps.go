package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		count, ok := wordMap[word]
		if ok {
			count++
			wordMap[word] = count
		} else {
			wordMap[word] = 1
		}
	}
	fmt.Println(words)
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
