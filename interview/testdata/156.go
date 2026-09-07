package main

import (
	"fmt"
)

type A interface{ F() int }
type B interface{ F() int }
type C interface {
	A
	B
}
type T struct{}

func (T) F() int { return 1 }
func main() {
	var i C = T{}
	fmt.Println(i.F())
}
