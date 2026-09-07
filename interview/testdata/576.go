package main

import (
	"errors"
	"fmt"
)

var sentinel = errors.New("target")

type E struct{}

func (E) Error() string   { return "E" }
func (E) Is(t error) bool { return t == sentinel }
func main() {
	fmt.Println(errors.Is(E{}, sentinel))
}
