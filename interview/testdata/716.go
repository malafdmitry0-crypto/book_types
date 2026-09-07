package main

import (
	"errors"
	"fmt"
)

func main() {
	a := errors.New("a")
	b := errors.New("b")
	e := errors.Join(nil, a, b)
	fmt.Println(errors.Is(e, a), errors.Is(e, b), errors.Join(nil, nil) == nil)
}
