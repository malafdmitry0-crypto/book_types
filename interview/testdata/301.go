package main

import (
	"fmt"
)

func max[T ~int | ~int32](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func main() {
	var n int32 = 2
	fmt.Printf("%T\n", max(n, 3))
}
