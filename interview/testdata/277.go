package main

import (
	"fmt"
)

type Box[T any] struct{ V T }

func main() {
	var a any = 7
	b := Box[any]{a}
	fmt.Printf("%T %T\n", b, b.V)
}
