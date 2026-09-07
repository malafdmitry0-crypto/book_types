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
	s := Reduce([]int{1, 2, 3}, "", func(a string, v int) string { return a + fmt.Sprint(v) })
	fmt.Printf("%T %v\n", s, s)
}
