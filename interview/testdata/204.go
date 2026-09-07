package main

import (
	"fmt"
)

func main() {
	a := map[int]int{1: 2}
	b := a
	b[1] = 9
	fmt.Println(a[1])
}
