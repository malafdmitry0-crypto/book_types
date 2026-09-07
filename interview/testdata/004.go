package main

import (
	"fmt"
)

func main() {
	const n = 1 << 100
	var x any = n
	fmt.Println(x)
}
