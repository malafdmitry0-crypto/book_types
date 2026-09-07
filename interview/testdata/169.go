package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	a = append(a[:1], a[2:]...)
	fmt.Println(a)
}
