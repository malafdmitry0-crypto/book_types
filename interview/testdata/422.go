package main

import (
	"fmt"
)

type A [2]int

func (a A) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func main() {
	a := A{2, 1}
	a.Swap(0, 1)
	fmt.Println(a)
}
