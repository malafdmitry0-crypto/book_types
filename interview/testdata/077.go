package main

import (
	"fmt"
)

type A = int

func main() {
	a := []A{1}
	var b []int = a
	fmt.Println(b)
}
