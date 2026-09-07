package main

import (
	"fmt"
)

type E struct{}

func (*E) Error() string { return "E" }
func main() {
	var p *E
	var e error = p
	fmt.Println(p == nil, e == nil)
}
