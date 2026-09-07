package main

import (
	"fmt"
	"io"
)

type S struct{}

func (S) Read([]byte) (int, error) {
	fmt.Println("read")
	return 0, io.EOF
}
func (S) WriteTo(io.Writer) (int64, error) {
	fmt.Println("source")
	return 7, nil
}

type D struct{}

func (D) Write(p []byte) (int, error) { return len(p), nil }
func (D) ReadFrom(io.Reader) (int64, error) {
	fmt.Println("destination")
	return 9, nil
}
func main() {
	s := struct{ io.Reader }{S{}}
	n, _ := io.Copy(D{}, s)
	fmt.Println(n)
}
