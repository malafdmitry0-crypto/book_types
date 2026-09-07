package main

import (
	"errors"
	"fmt"
)

type E struct{}

func (*E) Error() string { return "E" }
func main() {
	var p *E
	var err error = p
	var found *E
	ok := errors.As(err, &found)
	fmt.Println(ok, found == nil, err == nil)
}
