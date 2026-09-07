package main

import (
	"errors"
	"fmt"
)

type E struct{}

func (*E) Error() string { return "E" }
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	err := &E{}
	var found *E
	errors.As(err, found)
}
