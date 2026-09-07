package main

import (
	"fmt"
)

func first[T any](a, b T) T { return a }
func main() {
	var p *int
	fmt.Println(first(p, nil) == nil)
}
