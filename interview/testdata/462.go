package main

import (
	"fmt"
)

func compose(f, g func(int) int) func(int) int { return func(x int) int { return f(g(x)) } }
func main() {
	n := 0
	f := func(x int) int {
		n++
		return x
	}
	c := compose(f, f)
	fmt.Print(n, " ")
	fmt.Println(c(1), n)
}
