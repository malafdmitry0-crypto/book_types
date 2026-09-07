package main

import (
	"fmt"
)

type A = struct{ X int }

func main() {
	var a A
	var b struct{ X int }
	a = b
	fmt.Printf("%T\n", a)
}
