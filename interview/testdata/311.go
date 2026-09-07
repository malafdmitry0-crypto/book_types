package main

import (
	"fmt"
)

func convert[To, From ~int32 | ~int64](v From) To { return To(v) }
func main() {
	fmt.Printf("%T %v\n", convert[int64](int32(7)), convert[int64](int32(7)))
}
