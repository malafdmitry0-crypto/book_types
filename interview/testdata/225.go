package main

import (
	"fmt"
)

func f(x ...any) { fmt.Printf("%d %T\n", len(x), x[0]) }
func main() {
	a := []int{1, 2}
	f(a)
}
