package main

import (
	"fmt"
)

func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
func main() {
	a, b := counter(), counter()
	fmt.Println(a(), b(), a())
}
