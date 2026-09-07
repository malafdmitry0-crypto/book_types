package main

import (
	"fmt"
	"io"
)

type R struct{}

func (R) Read(p []byte) (int, error) { return copy(p, "ok"), io.EOF }
func main() {
	b, e := io.ReadAll(R{})
	fmt.Println(string(b), e == nil)
}
