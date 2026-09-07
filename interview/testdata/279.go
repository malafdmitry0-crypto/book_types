package main

import (
	"fmt"
)

func id[T any](v T) T { return v }
func main() {
	f := id[int]
	fmt.Printf("%T %v\n", f, f(9))
}
