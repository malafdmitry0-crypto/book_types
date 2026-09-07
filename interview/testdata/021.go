package main

import (
	"fmt"
)

func main() {
	var n int16 = 255
	fmt.Println(int8(n), int16(int8(n)))
}
