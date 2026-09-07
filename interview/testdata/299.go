package main

import (
	"fmt"
)

func empty[S ~[]E, E any]() S { return nil }
func main() {
	fmt.Println(empty[[]int]() == nil)
}
