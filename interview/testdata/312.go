package main

import (
	"fmt"
)

func pair[A, B any](a A, b B) A { return a }
func main() {
	fmt.Println(pair[_, int](1, 2))
}
