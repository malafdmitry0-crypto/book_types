package main

import (
	"fmt"
)

func main() {
	n := 0
	f := func() bool {
		n++
		return true
	}
	fmt.Println(false && f(), true || f(), n)
}
