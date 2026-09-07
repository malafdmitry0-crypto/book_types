package main

import (
	"fmt"
)

func Pick[T any](a, b T) T { return a }
func main() {
	fmt.Printf("%T %v", Pick[int8](3, 4), Pick[int8](3, 4))
}
