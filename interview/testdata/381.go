package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n := 0
	f := func() int {
		n++
		return 1
	}
	_ = unsafe.Sizeof(f())
	fmt.Println(n)
}
