package main

import (
	"fmt"
)

func first[T any](a, b T) T { return a }
func main() {
	var a uint8 = 1
	fmt.Println(first(a, 256))
}
