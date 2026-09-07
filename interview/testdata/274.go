package main

import (
	"fmt"
)

type Box[T any] struct{ V T }

func (b Box[T]) Map[R any](f func(T) R) R { return f(b.V) }
func main() {
	fmt.Println(Box[int]{})
}
