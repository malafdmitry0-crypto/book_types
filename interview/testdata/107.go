package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	a := [1]any{[]int{}}
	fmt.Println(a == a)
}
