package main

import (
	"fmt"
)

type A int
type B int

func main() {
	var a A = 3
	b := (*B)(&a)
	*b = 7
	fmt.Println(a)
}
