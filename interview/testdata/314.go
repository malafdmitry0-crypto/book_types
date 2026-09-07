package main

import (
	"fmt"
)

type N []int

func clone[E any](s []E) []E { return append([]E(nil), s...) }
func main() {
	x := clone(N{1, 2})
	fmt.Printf("%T\n", x)
}
