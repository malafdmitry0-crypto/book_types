package main

import (
	"fmt"
)

func Chain(fs []func(int) int) func(int) int {
	return func(x int) int {
		for _, f := range fs {
			x = f(x)
		}
		return x
	}
}
func main() {
	fs := []func(int) int{func(x int) int { return x + 1 }}
	f := Chain(fs)
	fs[0] = func(x int) int { return x * 10 }
	fmt.Println(f(2))
}
