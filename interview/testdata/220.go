package main

import (
	"fmt"
)

func main() {
	m := map[int]int{1: 0}
	a, okA := m[1]
	b, okB := m[2]
	fmt.Println(a, b, okA, okB)
}
