package main

import (
	"fmt"
)

func main() {
	var n int64 = 16777217
	fmt.Println(float64(n), float64(float32(n)))
}
