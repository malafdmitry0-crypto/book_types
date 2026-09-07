package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	n := copy(a[1:], a)
	m := copy(a, a[2:])
	fmt.Println(n, m, a)
}
