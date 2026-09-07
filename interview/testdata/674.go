package main

import (
	"fmt"
)

func Pair[A, B any](a A, b B) struct {
	A A
	B B
} {
	return struct {
		A A
		B B
	}{a, b}
}
func main() {
	fmt.Printf("%T", Pair[int](1, "x"))
}
