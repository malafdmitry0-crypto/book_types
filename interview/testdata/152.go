package main

import (
	"fmt"
)

type A struct{ X int }
type O struct{ A }

func main() {
	o := O{A: A{X: 1}}
	fmt.Println(o.X)
}
