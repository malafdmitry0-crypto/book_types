package main

import (
	"fmt"
)

type Option[T any] struct {
	v  T
	ok bool
}

func MapOption[A, B any](o Option[A], f func(A) B) Option[B] {
	if !o.ok {
		return Option[B]{}
	}
	return Option[B]{f(o.v), true}
}
func main() {
	n := 0
	var o Option[int]
	r := MapOption(o, func(x int) string {
		n++
		return "x"
	})
	fmt.Println(r.ok, n)
}
