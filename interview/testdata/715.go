package main

import (
	"errors"
	"fmt"
)

func main() {
	a := errors.New("x")
	b := errors.New("x")
	fmt.Println(a == b, errors.Is(a, b))
}
