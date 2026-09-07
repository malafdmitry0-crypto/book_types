package main

import (
	"fmt"
)

func main() {
	m := map[int][]int{1: {2}}
	m[1][0] = 9
	fmt.Println(m[1])
}
