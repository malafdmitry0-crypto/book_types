package main

import (
	"fmt"
)

func main() {
	var a any = []int{}
	var b any = map[int]int{}
	fmt.Println(a == b)
}
