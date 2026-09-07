package main

import (
	"fmt"
)

type F func(int) int

func main() {
	var f F = func(x int) int { return x + 1 }
	fmt.Println(f(4))
}
