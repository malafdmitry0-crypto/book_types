package main

import (
	"fmt"
)

type F func(int) int

func (f F) Apply(n int) int { return f(n) }
func (f F) Twice(n int) int { return f(f(n)) }
func main() {
	f := F(func(n int) int { return n + 1 })
	var a interface{ Apply(int) int } = f
	var b interface{ Twice(int) int } = f
	fmt.Println(a.Apply(2), b.Twice(2))
}
