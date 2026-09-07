package main

import (
	"fmt"
	"iter"
	"slices"
)

func Once[E any](s []E) iter.Seq[E] {
	i := 0
	return func(yield func(E) bool) {
		for i < len(s) {
			v := s[i]
			i++
			if !yield(v) {
				return
			}
		}
	}
}
func main() {
	s := Once([]int{1, 2, 3})
	for v := range s {
		fmt.Println(v)
		break
	}
	fmt.Println(slices.Collect(s))
}
