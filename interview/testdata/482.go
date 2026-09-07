package main

import (
	"fmt"
)

type Names []string

func Map[Out ~[]R, E, R any](s []E, f func(E) R) Out {
	out := make(Out, len(s))
	for i, v := range s {
		out[i] = f(v)
	}
	return out
}
func main() {
	out := Map[Names]([]int{2, 3}, func(n int) string { return fmt.Sprint(n) })
	fmt.Printf("%T %v\n", out, out)
}
