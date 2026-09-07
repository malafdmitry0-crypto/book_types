package main

import (
	"fmt"
)

type Strategy[T any] interface{ Apply(T) T }
type Counter struct{ N int }

func (c *Counter) Apply(x int) int {
	c.N++
	return x
}
func Run[T ~int, S Strategy[T]](s S, xs []T) T {
	var sum T
	for _, x := range xs {
		sum += s.Apply(x)
	}
	return sum
}
func main() {
	c := &Counter{}
	fmt.Println(Run[int](c, []int{2, 3}), c.N)
}
