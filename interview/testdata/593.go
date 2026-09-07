package main

import (
	"fmt"
	"iter"
	"slices"
)

type Numbers []int

func (s *Numbers) All() iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, v := range *s {
			if !yield(v) {
				return
			}
		}
	}
}
func main() {
	s := Numbers{1}
	seq := s.All()
	s = append(s, 2)
	fmt.Println(slices.Collect(seq))
}
