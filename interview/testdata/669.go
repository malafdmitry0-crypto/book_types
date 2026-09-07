package main

import (
	"fmt"
)

func Pick[T any](a, b T) T { return a }
func main() {
	fmt.Printf("%T", Pick(1, 2.5))
}
