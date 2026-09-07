package main

import (
	"errors"
	"fmt"
)

type E struct{}

func (*E) Error() string   { return "E" }
func (*E) Temporary() bool { return true }
func main() {
	err := fmt.Errorf("layer: %w", &E{})
	var found interface{ Temporary() bool }
	ok := errors.As(err, &found)
	fmt.Println(ok, found.Temporary())
}
