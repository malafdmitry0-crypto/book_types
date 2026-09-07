package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.Join(errors.New("a"), errors.New("b"))
	fmt.Println(errors.Unwrap(e) == nil)
}
