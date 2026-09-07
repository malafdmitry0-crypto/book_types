package main

import (
	"fmt"
)

func first[T any](a, b T) T { return a }
func main() {
	fmt.Printf("%T %v\n", first(1, 2.5), first(1, 2.5))
}
