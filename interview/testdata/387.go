package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [2]uint32{1, 2}
	p := (*uint32)(unsafe.Add(unsafe.Pointer(&a[0]), unsafe.Sizeof(a[0])))
	*p = 9
	fmt.Println(a)
}
