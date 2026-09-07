package main

import (
	"fmt"
)

func curry(a int) func(int) int { return func(b int) int { return a + b } }
func main() {
	add2 := curry(2)
	add5 := curry(5)
	fmt.Println(add2(3), add5(3))
}
