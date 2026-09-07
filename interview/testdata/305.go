package main

import (
	"fmt"
)

func first[T any](a, b T) T { return a }
func main() {
	fmt.Printf("%T\n", first[any](int32(1), "x"))
}
