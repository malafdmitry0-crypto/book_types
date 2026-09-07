package main

import (
	"fmt"
)

func main() {
	var a, b uint8 = 200, 100
	fmt.Println(uint16(a+b), uint16(a)+uint16(b))
}
