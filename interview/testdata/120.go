package main

import (
	"fmt"
	"slices"
)

func main() {
	var a []int
	b := []int{}
	fmt.Println(slices.Equal(a, b))
}
