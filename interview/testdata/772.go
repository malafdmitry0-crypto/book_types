package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 1, 2, 1}
	b := slices.Compact(a)
	fmt.Println(b, a)
}
