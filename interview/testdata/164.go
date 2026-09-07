package main

import (
	"fmt"
)

func f(s []int) { s = append(s, 7) }
func main() {
	a := make([]int, 1, 2)
	f(a)
	fmt.Println(len(a), a[:2])
}
