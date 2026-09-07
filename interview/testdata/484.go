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
	var in []int
	out := Map(in, func(n int) string { return fmt.Sprint(n) })
	fmt.Println(out == nil, len(out))
}
