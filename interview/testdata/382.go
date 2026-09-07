package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a any = [100]int{}
	var b any = byte(1)
	fmt.Println(unsafe.Sizeof(a) == unsafe.Sizeof(b))
}
