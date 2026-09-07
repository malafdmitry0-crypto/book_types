package main

import (
	"fmt"
)

func main() {
	x := 1
	f := func() int { return x }
	x = 2
	fmt.Println(f())
}
