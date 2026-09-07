package main

import (
	"fmt"
)

func eq[T comparable](a, b T) bool { return a == b }
func main() {
	fmt.Println(eq[any](1, 1))
}
