package main

import (
	"golang.org/x/tour/reader"
)

/*
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = byte('A')
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
