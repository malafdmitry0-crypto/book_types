package main

import (
	"fmt"
)

type C interface{ Get() int }
type Counter int

func (c Counter) Get() int { return int(c) }
func read[T C](v T) int    { return v.Get() }
func main() {
	var c C = Counter(3)
	fmt.Println(read[C](c))
}
