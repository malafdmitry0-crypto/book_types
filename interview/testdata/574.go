package main

import (
	"errors"
	"fmt"
)

func main() {
	a, b := fmt.Errorf("a"), fmt.Errorf("b")
	err := errors.Join(a, b)
	fmt.Println(errors.Is(err, a), errors.Is(err, b), errors.Unwrap(err) == nil)
}
