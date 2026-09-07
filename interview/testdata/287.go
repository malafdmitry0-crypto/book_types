package main

import (
	"fmt"
)

type A struct{ X int }
type B struct {
	X int
	Y int
}

func f[T A | B](v T) int { return v.X }
func main() {
	fmt.Println(f(A{1}))
}
