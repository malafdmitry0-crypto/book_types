package main

import (
	"errors"
	"fmt"
)

func main() {
	base := errors.New("x")
	a := fmt.Errorf("wrap: %v", base)
	b := fmt.Errorf("wrap: %w", base)
	fmt.Println(errors.Is(a, base), errors.Is(b, base))
}
