package main

import (
	"fmt"
)

func main() {
	n := 0
	var fs []func() int
	for _, v := range []int{1, 2} {
		n += v
		snapshot := n
		fs = append(fs, func() int { return snapshot })
	}
	fmt.Println(fs[0](), fs[1]())
}
