package main

import (
	"fmt"
	"iter"
	"slices"
)

func Map[A, B any](s iter.Seq[A], f func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for x := range s {
			if !yield(f(x)) {
				return
			}
		}
	}
}
func main() {
	n := 0
	seq := Map(slices.Values([]int{1, 2}), func(x int) int {
		n++
		return x
	})
	slices.Collect(seq)
	slices.Collect(seq)
	fmt.Println(n)
}
