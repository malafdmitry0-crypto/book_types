package main

import (
	"fmt"
)

func f(x ...int) { x[0] = 9 }
func main() {
	a, b := 1, 2
	f(a, b)
	fmt.Println(a, b)
}
