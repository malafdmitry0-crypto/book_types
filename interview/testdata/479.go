package main

import (
	"fmt"
)

func compose(f, g func(any) any) func(any) any { return func(x any) any { return f(g(x)) } }
func main() {
	fmt.Println(compose(func(int) string { return "x" }, func(string) int { return 1 })(1))
}
