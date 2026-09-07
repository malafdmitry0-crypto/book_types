package main

import (
	"fmt"
)

func Filter[S ~[]E, E any](s S, p func(E) bool) S {
	out := s[:0]
	for _, v := range s {
		if p(v) {
			out = append(out, v)
		}
	}
	return out
}
func main() {
	s := []int{1, 2, 3}
	out := Filter(s, func(n int) bool { return n > 1 })
	fmt.Println(out, s)
}
