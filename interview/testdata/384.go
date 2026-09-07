package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p *int
	s := unsafe.Slice(p, 0)
	fmt.Println(s == nil)
}
