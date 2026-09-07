package main

import (
	"fmt"
)

func orElse(v int, ok bool, f func() int) int {
	if ok {
		return v
	}
	return f()
}
func main() {
	n := 0
	fallback := func() int {
		n++
		return 9
	}
	fmt.Println(orElse(3, true, fallback), n)
}
