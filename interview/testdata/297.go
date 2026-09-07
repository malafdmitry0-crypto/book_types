package main

import (
	"fmt"
)

func first[T ~string | ~[]byte](v T) byte { return v[0] }
func main() {
	fmt.Printf("%T %v\n", first("go"), first([]byte{9}))
}
