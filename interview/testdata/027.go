package main

import (
	"fmt"
)

func main() {
	var n uint8 = 0
	fmt.Println(^n, ^uint16(n), uint16(^n))
}
