package main

import (
	"fmt"
)

func id[T any](v T) T { return v }
func main() {
	var f func(int) int = id
	fmt.Println(f(9))
}
