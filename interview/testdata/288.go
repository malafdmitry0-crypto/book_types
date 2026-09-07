package main

import (
	"fmt"
)

type N int

func (n N) Double() int                      { return int(n * 2) }
func f[T interface{ Double() int }](v T) int { return v.Double() }
func main() {
	fmt.Println(f(N(3)))
}
