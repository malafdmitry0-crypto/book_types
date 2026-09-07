package main

import (
	"fmt"
)

func Zero[T any]() T {
	var z T
	return z
}
func main() {
	var x int = Zero()
	fmt.Println(x)
}
