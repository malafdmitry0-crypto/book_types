package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s []int
	fmt.Println(unsafe.SliceData(s) == nil)
}
