package main

import (
	"fmt"
	"iter"
	"slices"
)

func FilterSeq[E any](s iter.Seq[E], p func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s {
			if p(v) && !yield(v) {
				return
			}
		}
	}
}
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
	calls := 0
	s := MapSeq(FilterSeq(slices.Values([]int{1, 2, 3}), func(n int) bool { return n%2 == 1 }), func(n int) int {
		calls++
		return n * 10
	})
	fmt.Println(slices.Collect(s), calls)
}
