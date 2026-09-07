package main

import (
	"fmt"
)

func main() {
	n := 1
	f := func(v int) { fmt.Println(v, n) }
	arg := func() int {
		old := n
		n = 9
		return old
	}
	f(arg())
}
