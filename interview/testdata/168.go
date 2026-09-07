package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	copy(a, a[1:])
	fmt.Println(a)
}
