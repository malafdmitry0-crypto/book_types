package main

import (
	"fmt"
)

type F func(int) int

func (f F) Then(g func(int) int) F { return func(n int) int { return g(f(n)) } }
func main() {
	f := F(func(n int) int { return n * 2 })
	g := f.Then(func(n int) int { return n + 1 })
	fmt.Printf("%T %v\n", g, g(3))
}
