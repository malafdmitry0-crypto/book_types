package main

import (
	"fmt"
	"io"
)

type RF func([]byte) (int, error)

func (f RF) Read(p []byte) (int, error) { return f(p) }
func main() {
	r := RF(func(p []byte) (int, error) {
		p[0] = 'x'
		return 1, io.EOF
	})
	b, err := io.ReadAll(r)
	fmt.Printf("%s %v\n", b, err)
}
