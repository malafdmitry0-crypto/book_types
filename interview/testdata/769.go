package main

import (
	"fmt"
	"slices"
)

func main() {
	a := [][]int{{1}}
	b := slices.Clone(a)
	b[0][0] = 7
	b = append(b, []int{8})
	fmt.Println(a, b)
}
