package main

import (
	"fmt"
)

func Equal[T comparable](a, b T) bool { return a == b }
func main() {
	fmt.Println(Equal[any](1, int64(1)))
}
