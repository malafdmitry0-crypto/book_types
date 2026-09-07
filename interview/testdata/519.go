package main

import (
	"fmt"
)

type Op[E any] interface{ Apply(E, E) E }
type Plus struct{}

func (Plus) Apply(a, b int) int { return a + b }

type Mul struct{}

func (Mul) Apply(a, b int) int { return a * b }
func main() {
	ops := []Op[int]{Plus{}, Mul{}}
	a := 2
	for _, o := range ops {
		a = o.Apply(a, 3)
	}
	fmt.Println(a)
}
