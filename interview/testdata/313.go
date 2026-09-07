package main

import (
	"fmt"
)

type N []int

func clone[S ~[]E, E any](s S) S { return append(S(nil), s...) }
func main() {
	x := clone(N{1, 2})
	fmt.Printf("%T %v\n", x, x)
}
