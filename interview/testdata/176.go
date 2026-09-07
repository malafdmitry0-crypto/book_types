package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 2}
	b := slices.Clip(a)
	b[0] = 9
	fmt.Println(a, cap(b) == len(b))
}
