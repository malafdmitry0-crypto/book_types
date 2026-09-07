package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := slices.Delete(a, 1, 3)
	fmt.Println(b, a)
}
