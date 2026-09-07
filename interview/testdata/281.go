package main

import (
	"fmt"
)

type N int

func f[T int](v T) T { return v }
func main() {
	fmt.Println(f(N(1)))
}
