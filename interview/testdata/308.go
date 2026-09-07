package main

import (
	"fmt"
)

func zero[T any]() T {
	var z T
	return z
}
func main() {
	var x int = zero()
	fmt.Println(x)
}
