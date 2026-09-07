package main

import (
	"fmt"
)

func f(x ...int) bool { return x == nil }
func main() {
	fmt.Println(f(), f([]int{}...))
}
