package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	a := make([]int, 1, 3)
	fmt.Println(a[2])
}
