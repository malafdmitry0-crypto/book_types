package main

import (
	"fmt"
)

func Map[E, R any](s []E, f func(E) R) []R {
	out := make([]R, len(s))
	for i, v := range s {
		out[i] = f(v)
	}
	return out
}
func main() {
	out := Map([]int{2, 3}, func(n int) string { return fmt.Sprint(n * n) })
	fmt.Printf("%T %v\n", out, out)
}
