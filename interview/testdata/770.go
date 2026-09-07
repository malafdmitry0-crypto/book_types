package main

import (
	"fmt"
	"slices"
)

func main() {
	var a []int
	b := slices.Clone(a)
	fmt.Println(b == nil, len(b))
}
