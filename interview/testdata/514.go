package main

import (
	"fmt"
)

type Mapper[A, B any] interface{ Map(A) B }
type MF[A, B any] func(A) B

func (f MF[A, B]) Map(a A) B { return f(a) }
func main() {
	var m Mapper[int, string] = MF[int, string](func(n int) string { return fmt.Sprint(n) })
	_, ok := m.(MF[int, string])
	fmt.Println(ok)
}
