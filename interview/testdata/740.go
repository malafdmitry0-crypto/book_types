package main

import (
	"fmt"
)

func Bind(f func(int, int) int, a int) func(int) int { return func(b int) int { return f(a, b) } }
func main() {
	k := 2
	f := Bind(func(a, b int) int { return a * b }, k)
	k = 9
	fmt.Println(f(3), k)
}
