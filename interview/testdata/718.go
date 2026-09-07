package main

import (
	"errors"
	"fmt"
)

type E struct{ Code int }

func (e *E) Error() string { return "E" }
func main() {
	var target *E
	err := fmt.Errorf("wrap: %w", &E{7})
	ok := errors.As(err, &target)
	fmt.Println(ok, target.Code)
}
