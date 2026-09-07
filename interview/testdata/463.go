package main

import (
	"fmt"
)

func bind(f func(int, int) int, a int) func(int) int { return func(b int) int { return f(a, b) } }
func main() {
	n := 2
	f := bind(func(a, b int) int { return a + b }, n)
	n = 9
	fmt.Println(f(3))
}
