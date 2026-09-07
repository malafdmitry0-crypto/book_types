package main

import (
	"fmt"
	"io"
)

type RF func([]byte) (int, error)

func (f RF) Read(p []byte) (int, error) { return f(p) }
func main() {
	n := 0
	r := RF(func(p []byte) (int, error) {
		if n == 2 {
			return 0, io.EOF
		}
		p[0] = byte(65 + n)
		n++
		return 1, nil
	})
	b, _ := io.ReadAll(r)
	fmt.Println(string(b), n)
}
