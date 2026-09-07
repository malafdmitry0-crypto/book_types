package main

import (
	"fmt"
)

func f(x ...int) { x[0] = 9 }
func main() {
	a := []int{1, 2}
	f(a...)
	fmt.Println(a)
}
