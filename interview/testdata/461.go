package main

import (
	"fmt"
)

func compose(f, g func(int) int) func(int) int { return func(x int) int { return f(g(x)) } }
func main() {
	inc := func(n int) int { return n + 1 }
	dbl := func(n int) int { return n * 2 }
	fmt.Println(compose(inc, dbl)(3), compose(dbl, inc)(3))
}
