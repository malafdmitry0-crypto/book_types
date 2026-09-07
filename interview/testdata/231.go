package main

import (
	"fmt"
)

func main() {
	x := 1
	defer func() { fmt.Println(x) }()
	x = 2
}
