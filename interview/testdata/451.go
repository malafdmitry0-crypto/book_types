package main

import (
	"fmt"
)

func main() {
	n := 0
	var fs []func() int
	for _, v := range []int{1, 2} {
		n += v
		fs = append(fs, func() int { return n })
	}
	fmt.Println(fs[0](), fs[1]())
}
