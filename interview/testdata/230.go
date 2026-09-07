package main

import (
	"fmt"
)

func main() {
	x := 1
	defer fmt.Println(x)
	x = 2
}
