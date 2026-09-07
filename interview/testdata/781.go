package main

import (
	"fmt"
)

type IDs []int

func Map[S ~[]B, A, B any](xs []A, f func(A) B) S {
	out := make(S, len(xs))
	for i, x := range xs {
		out[i] = f(x)
	}
	return out
}
func main() {
	x := Map[IDs]([]string{"go", "types"}, func(s string) int { return len(s) })
	fmt.Printf("%T %v", x, x)
}
