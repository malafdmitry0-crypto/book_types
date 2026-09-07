package main

import (
	"fmt"
)

type E struct{}

func (*E) Error() string { return "E" }
func main() {
	var e error = (*E)(nil)
	var a any = e
	fmt.Printf("%T %v\n", a, a == nil)
}
