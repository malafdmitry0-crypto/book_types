package main

import (
	"fmt"
)

func main() {
	var fs []func() int
	for _, v := range []int{2, 4, 6} {
		fs = append(fs, func() int { return v })
	}
	fmt.Println(fs[0](), fs[1](), fs[2]())
}
