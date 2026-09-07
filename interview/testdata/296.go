package main

import (
	"fmt"
)

func size[T ~string | ~[]int](v T) int { return len(v) }
func main() {
	fmt.Println(size("я"), size([]int{1, 2, 3}))
}
