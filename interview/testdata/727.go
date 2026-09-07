package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type W struct{}

func (W) Write([]byte) (int, error) { return 0, errors.New("stop") }
func main() {
	var b bytes.Buffer
	n, e := io.MultiWriter(W{}, &b).Write([]byte("go"))
	fmt.Println(n, e != nil, b.Len())
}
