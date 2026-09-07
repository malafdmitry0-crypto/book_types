package main

import (
	"fmt"
)

type A int

func main() {
	a := []A{1}
	fmt.Println([]int(a))
}
