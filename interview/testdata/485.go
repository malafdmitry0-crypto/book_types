package main

import (
	"fmt"
)

type Numbers []int

func Filter[S ~[]E, E any](s S, p func(E) bool) S {
	out := make(S, 0, len(s))
	for _, v := range s {
		if p(v) {
			out = append(out, v)
		}
	}
	return out
}
func main() {
	out := Filter(Numbers{1, 2, 3}, func(n int) bool { return n > 1 })
	fmt.Printf("%T %v\n", out, out)
}
