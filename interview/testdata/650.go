package main

import (
	"fmt"
)

func main() {
	var a any
	var b any = []int{1}
	fmt.Println(a == b)
}
