package main

import (
	"fmt"
)

func main() {
	var a int8 = -8
	fmt.Println(a>>2, uint8(a)>>2)
}
