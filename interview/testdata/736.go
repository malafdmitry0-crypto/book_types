package main

import (
	"fmt"
)

func Twice(f func(int) int) func(int) int { return func(x int) int { return f(f(x)) } }
func main() {
	n := 0
	makeF := func() func(int) int {
		n++
		return func(x int) int { return x + n }
	}
	f := Twice(makeF())
	fmt.Println(f(1), n)
}
