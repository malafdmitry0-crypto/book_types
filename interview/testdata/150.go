package main

import (
	"fmt"
)

type A struct{ X int }
type O struct {
	A
	X int
}

func main() {
	o := O{A: A{X: 1}, X: 2}
	fmt.Println(o.X, o.A.X)
}
