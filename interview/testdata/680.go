package main

import (
	"fmt"
)

type Box[T any] struct{ V T }

func main() {
	var x Box[any] = Box[int]{1}
	fmt.Println(x)
}
