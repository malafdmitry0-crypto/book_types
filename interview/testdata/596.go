package main

import (
	"fmt"
	"iter"
	"slices"
)

func MapSeq[A, B any](s iter.Seq[A], f func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for v := range s {
			if !yield(f(v)) {
				return
			}
		}
	}
}
func main() {
	s := MapSeq(slices.Values([]int{1, 2}), func(n int) any { return n })
	out := slices.Collect(s)
	fmt.Printf("%T %T\n", out, out[0])
}
