package main

import (
	//"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(s []byte) (int, error) {
	for i := range s {
		s[i] = 65
	}
	return len(s), nil
}

func main() {
	reader.Validate(MyReader{})
}
