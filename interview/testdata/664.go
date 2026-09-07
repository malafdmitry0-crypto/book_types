package main

import (
	"fmt"
)

func Equal[T comparable](a, b T) bool { return a == b }
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	fmt.Println(Equal[any]([]int{1}, []int{1}))
}
