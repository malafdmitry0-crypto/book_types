package main

import (
	"fmt"
)

func main() {
	a := []int{1}
	a[0], a[0] = 2, 3
	fmt.Println(a)
}
