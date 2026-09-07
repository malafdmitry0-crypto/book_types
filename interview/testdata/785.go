package main

import (
	"fmt"
)

func Filter[T any](xs []T, p func(T) bool) []T {
	out := xs[:0]
	for _, x := range xs {
		if p(x) {
			out = append(out, x)
		}
	}
	return out
}
func main() {
	a := []int{1, 2, 3, 4}
	b := Filter(a, func(x int) bool { return x%2 == 0 })
	fmt.Println(a, b)
}
