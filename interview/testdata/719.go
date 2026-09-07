package main

import (
	"errors"
	"fmt"
)

type E struct{}

func (*E) Error() string { return "E" }
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var target *E
	errors.As(&E{}, target)
}
