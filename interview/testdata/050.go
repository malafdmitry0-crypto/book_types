package main

import (
	"fmt"
)

type T int
type A = T

func (A) F() int { return 7 }
func main() {
	var x T
	fmt.Println(x.F())
}
