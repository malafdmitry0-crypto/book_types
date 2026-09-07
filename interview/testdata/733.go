package main

import (
	"fmt"
	"io"
	"strings"
)

type Count struct {
	io.Reader
	N int
}

func (c Count) Read(p []byte) (int, error) {
	n, e := c.Reader.Read(p)
	c.N += n
	return n, e
}
func main() {
	c := Count{Reader: strings.NewReader("abc")}
	io.ReadAll(c)
	fmt.Println(c.N)
}
