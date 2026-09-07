package main

import (
	"fmt"
	"slices"
)

func main() {
	x := 1
	a := []*int{&x}
	b := slices.Clone(a)
	*b[0] = 9
	b[0] = nil
	fmt.Println(*a[0], b[0] == nil)
}
