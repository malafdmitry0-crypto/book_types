package main

import (
	"fmt"
)

func Fold[A, B any](xs []A, z B, f func(B, A) B) B {
	for _, x := range xs {
		z = f(z, x)
	}
	return z
}
func main() {
	v := Fold([]int{2, 3}, "n", func(s string, x int) string { return fmt.Sprint(s, x) })
	fmt.Println(v)
}
