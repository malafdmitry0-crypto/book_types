package main

import (
	"fmt"
)

func main() {
	var f func(x int) (y int) = func(z int) int { return z }
	var g func(int) int = f
	fmt.Println(g(8))
}
