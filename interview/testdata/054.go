package main

import (
	"fmt"
)

type A []int

func main() {
	a := A{1}
	b := []int(a)
	b[0] = 9
	fmt.Println(a)
}
