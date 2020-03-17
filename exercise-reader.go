package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func main() {
	reader.Validate(MyReader{})
}
