package main

import (
	"fmt"
)

type F func(int) int
type G func(int) int

func main() {
	f := F(func(x int) int { return x + 1 })
	g := G(f)
	fmt.Println(g(4))
}
