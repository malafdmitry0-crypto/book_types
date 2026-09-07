package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	s := make([]int, 1, 3)
	_ = (*[2]int)(s)
}
