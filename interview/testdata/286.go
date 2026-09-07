package main

import (
	"fmt"
)

func f[T ~int | ~string](v T) T { return v - v }
func main() {
	fmt.Println(f(3))
}
