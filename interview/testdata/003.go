package main

import (
	"fmt"
)

func main() {
	const n = 1 << 100
	fmt.Println(n >> 99)
}
