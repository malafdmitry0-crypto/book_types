package main

import (
	"fmt"
)

func main() {
	s := []int{1}
	var a any = s
	s[0] = 2
	fmt.Println(a.([]int))
}
