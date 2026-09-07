package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	s := append(a[:1], a[1:3]...)
	fmt.Println(a, s)
}
