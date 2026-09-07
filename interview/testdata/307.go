package main

import (
	"fmt"
)

func id[T any](v T) T { return v }
func main() {
	fmt.Printf("%T\n", id((*int)(nil)))
}
