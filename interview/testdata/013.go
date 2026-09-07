package main

import (
	"fmt"
)

func main() {
	var s uint = 3
	var x uint64 = 1 << s
	fmt.Printf("%T %v\n", x, x)
}
