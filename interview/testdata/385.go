package main

import (
	"fmt"
	"unsafe"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var p *int
	_ = unsafe.Slice(p, 1)
}
