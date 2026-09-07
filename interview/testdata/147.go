package main

import (
	"fmt"
)

type A struct{}

func (A) F() string { return "A" }

type B struct{}

func (B) F() string { return "B" }

type O struct {
	A
	B
}

func main() {
	fmt.Println(O{}.A.F(), O{}.B.F())
}
