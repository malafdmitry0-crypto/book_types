package main

import (
	"fmt"
)

func main() {
	x := 1
	y := x
	f := func() int { return y }
	x = 2
	fmt.Println(f(), x)
}
