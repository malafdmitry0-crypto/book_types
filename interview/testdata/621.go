package main

import (
	"fmt"
)

func main() {
	x := 1
	if x := x + 2; x > 0 {
		x++
		fmt.Print(x, " ")
	}
	fmt.Println(x)
}
