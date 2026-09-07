package main

import (
	"fmt"
)

func main() {
	const a float32 = 16777216
	const b float32 = a + 1
	fmt.Println(a == b)
}
