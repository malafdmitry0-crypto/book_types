package main

import (
	"fmt"
)

func compose[A, B, C any](f func(B) C, g func(A) B) func(A) C { return func(a A) C { return f(g(a)) } }
func main() {
	f := compose(func(n int) string { return fmt.Sprint(n) }, func(s string) int { return len(s) })
	fmt.Printf("%T %v\n", f, f("go"))
}
