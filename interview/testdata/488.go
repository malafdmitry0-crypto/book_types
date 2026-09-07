package main

import (
	"fmt"
)

func Reduce[E, A any](s []E, a A, f func(A, E) A) A {
	for _, v := range s {
		a = f(a, v)
	}
	return a
}
func main() {
	fmt.Printf("%T\n", Reduce([]int{1}, int64(0), func(a int64, v int) int64 { return a + int64(v) }))
}
