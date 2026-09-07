package main

import (
	"errors"
	"fmt"
)

type E struct{}

func (*E) Error() string { return "E" }
func main() {
	var p *E
	var e error = p
	j := errors.Join(e)
	fmt.Println(j == nil, errors.Is(j, e))
}
