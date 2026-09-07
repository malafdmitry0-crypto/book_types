package main

import (
	"fmt"
)

func main() {
	var a []int
	b := make([]int, 0)
	pa := (*[0]int)(a)
	pb := (*[0]int)(b)
	fmt.Println(pa == nil, pb == nil, [0]int(a) == [0]int(b))
}
