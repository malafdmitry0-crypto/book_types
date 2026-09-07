package main

import (
	"fmt"
)

func main() {
	i := 0
	a := []int{4, 5}
	i, a[i] = 1, 9
	fmt.Println(i, a)
}
