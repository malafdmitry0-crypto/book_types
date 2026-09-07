package main

import (
	"fmt"
)

func Compose[A, B, C any](f func(A) B, g func(B) C) func(A) C { return func(x A) C { return g(f(x)) } }
func main() {
	f := Compose(func(s string) int { return len(s) }, func(n int64) bool { return n > 0 })
	fmt.Println(f("go"))
}
