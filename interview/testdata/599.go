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
	var failure error
	s := MapSeq(slices.Values([]int{1, -1, 3}), func(n int) int {
		if n < 0 {
			failure = fmt.Errorf("bad")
		}
		return n
	})
	var out []int
	for n := range s {
		if failure != nil {
			break
		}
		out = append(out, n)
	}
	fmt.Println(out, failure != nil)
}
