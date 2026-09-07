package main

import (
	"fmt"
)

func Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C { return func(a A) C { return f(g(a)) } }
func main() {
	toText := func(n int) string { return fmt.Sprint(n) }
	toLen := func(s string) int { return len(s) }
	toBool := func(n int) bool { return n > 1 }
	f := Compose(toBool, Compose(toLen, toText))
	fmt.Printf("%T %v %v\n", f, f(7), f(42))
}
