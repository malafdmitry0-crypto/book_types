package main

import (
	"fmt"
)

func Less[T comparable](a, b T) bool { return a < b }
func main() {
	fmt.Println(Less(1, 2))
}
