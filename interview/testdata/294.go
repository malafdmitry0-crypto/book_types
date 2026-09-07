package main

import (
	"fmt"
)

func eq[T comparable](a, b T) bool { return a == b }
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	fmt.Println(eq[any]([]int{}, []int{}))
}
