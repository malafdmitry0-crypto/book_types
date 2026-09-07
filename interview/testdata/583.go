package main

import (
	"fmt"
	"iter"
	"slices"
)

func MapSeq[A, B any](src iter.Seq[A], f func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for v := range src {
			if !yield(f(v)) {
				return
			}
		}
	}
}
func main() {
	calls := 0
	s := MapSeq(slices.Values([]int{1, 2}), func(n int) int {
		calls++
		return n
	})
	_ = slices.Collect(s)
	_ = slices.Collect(s)
	fmt.Println(calls)
}
