package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var a any = []int{}
	fmt.Println(a == a)
}
