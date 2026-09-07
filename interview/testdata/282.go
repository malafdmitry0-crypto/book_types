package main

import (
	"fmt"
)

type N int

func f[T ~int](v T) T { return v + 1 }
func main() {
	fmt.Printf("%T %v\n", f(N(1)), f(N(1)))
}
