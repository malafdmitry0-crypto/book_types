package main

import (
	"fmt"
	"sort"
)

type A [2]int

func (a A) Len() int           { return len(a) }
func (a A) Less(i, j int) bool { return a[i] < a[j] }
func (a *A) Swap(i, j int)     { a[i], a[j] = a[j], a[i] }
func main() {
	a := A{2, 1}
	sort.Sort(&a)
	fmt.Println(a)
}
