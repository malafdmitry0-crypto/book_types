package main

import (
	"fmt"
)

func main() {
	a := make([]int, 1, 3)
	b := a[:3]
	fmt.Println(len(b), b)
}
