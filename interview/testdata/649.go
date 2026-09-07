package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var a any = []int(nil)
	var b any = []int(nil)
	fmt.Println(a == b)
}
