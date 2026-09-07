package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a, b := make([]byte, 1), make([]byte, 1000)
	fmt.Println(unsafe.Sizeof(a) == unsafe.Sizeof(b))
}
