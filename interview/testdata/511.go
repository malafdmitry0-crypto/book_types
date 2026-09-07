package main

import (
	"fmt"
)

type Op[E any] interface{ Apply(E, E) E }
type Plus struct{}

func (Plus) Apply(a, b int) int { return a + b }
func Fold[E any](s []E, a E, o Op[E]) E {
	for _, v := range s {
		a = o.Apply(a, v)
	}
	return a
}
func main() {
	fmt.Println(Fold([]int{2, 3}, 10, Plus{}))
}
