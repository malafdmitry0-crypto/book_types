package main

import (
	"fmt"
)

type N struct{ V int }

func (n *N) Add(x int) int {
	n.V += x
	return n.V
}
func main() {
	f := (*N).Add
	n := N{5}
	fmt.Println(f(&n, 3), n.V)
}
