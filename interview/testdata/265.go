package main

import (
	"fmt"
)

func zero[T any]() T {
	var z T
	return z
}
func main() {
	fmt.Printf("%T %v %q\n", zero[int](), zero[int](), zero[string]())
}
