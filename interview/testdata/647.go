package main

import (
	"fmt"
)

type A int
type B int

func main() {
	var a any = A(3)
	var b any = B(3)
	fmt.Println(a == b)
}
