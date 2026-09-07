package main

import (
	"fmt"
	"io"
	"strings"
)

type W struct{}

func (W) Write([]byte) (int, error) { return 1, nil }
func main() {
	n, e := io.Copy(W{}, strings.NewReader("abcd"))
	fmt.Println(n, e == io.ErrShortWrite)
}
