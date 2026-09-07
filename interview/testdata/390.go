package main

import (
	"fmt"
)

type Box[T any] struct{ V T }

func eq[T comparable](a, b T) bool { return a == b }
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	v := Box[any]{[]int{}}
	fmt.Println(eq(v, v))
}
