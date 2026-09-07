package main

import (
	"fmt"
)

func first[T ~[]int | ~[]string](v T) any { return v[0] }
func main() {
	fmt.Println(first([]int{1}))
}
