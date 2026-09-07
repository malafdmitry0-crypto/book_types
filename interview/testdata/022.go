package main

import (
	"fmt"
)

func main() {
	var n int8 = -1
	fmt.Println(uint16(n), uint16(uint8(n)))
}
