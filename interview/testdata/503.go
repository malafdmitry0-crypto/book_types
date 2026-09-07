package main

import (
	"fmt"
)

type Mapper[A, B any] interface{ Map(A) B }
type MF[A, B any] func(A) B

func (f MF[A, B]) Map(a A) B { return f(a) }

type Widen[A, B any] struct{ Inner Mapper[A, B] }

func (w Widen[A, B]) Map(a A) any { return w.Inner.Map(a) }
func main() {
	a := MF[int, string](func(n int) string { return fmt.Sprint(n) })
	var b Mapper[int, any] = Widen[int, string]{a}
	fmt.Printf("%T %v\n", b.Map(3), b.Map(3))
}
