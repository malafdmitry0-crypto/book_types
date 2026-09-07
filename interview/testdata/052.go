package main

import (
	"fmt"
)

type A [2]int
type B [2]int

func main() {
	a := A{1, 2}
	b := B(a)
	fmt.Printf("%T %T %v\n", a, b, a == A(b))
}
