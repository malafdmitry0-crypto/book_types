package main

import (
	"fmt"
)

func main() {
	s := []int{1}
	defer func() { fmt.Println(s) }()
	defer fmt.Println(s)
	s = []int{9}
}
