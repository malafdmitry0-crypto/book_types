package main

import (
	"fmt"
)

func main() {
	var a any = int32(1)
	var b any = int64(1)
	fmt.Println(a == b)
}
