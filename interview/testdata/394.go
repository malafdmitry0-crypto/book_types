package main

import (
	"fmt"
)

type N int

func (n N) Equal(m any) bool                          { return true }
func equal[T interface{ Equal(T) bool }](a, b T) bool { return a.Equal(b) }
func main() {
	fmt.Println(equal(N(1), N(1)))
}
