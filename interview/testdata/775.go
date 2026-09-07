package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	clear(a[1:])
	fmt.Println(a, len(a))
}
