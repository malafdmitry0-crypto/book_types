package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	n := copy(a[1:], a)
	fmt.Println(n, a)
}
