package main

import (
	"fmt"
)

type Step[E any] interface{ Add(E) }
type Acc struct{ N int }

func (a *Acc) Add(n int) { a.N += n }
func main() {
	a := &Acc{}
	var op Step[int] = a
	op.Add(2)
	op.Add(3)
	fmt.Println(a.N)
}
