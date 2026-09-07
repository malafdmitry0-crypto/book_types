package main

import (
	"fmt"
)

func main() {
	var fs []func() int
	for i := 0; i < 3; i++ {
		fs = append(fs, func() int { return i })
	}
	fmt.Println(fs[0](), fs[1](), fs[2]())
}
