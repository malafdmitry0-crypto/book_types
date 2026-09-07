package main

import (
	"fmt"
)

type N int

func (N) F() int                        { return 1 }
func f[T interface{ F() int }](v T) int { return (&v).F() }
func main() {
	fmt.Println(f(N(1)))
}
