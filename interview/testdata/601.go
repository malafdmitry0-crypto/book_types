package main

import (
	"fmt"
)

func main() {
	x := 2
	defer fmt.Println(x)
	defer func() { fmt.Println(x) }()
	x = 9
}
