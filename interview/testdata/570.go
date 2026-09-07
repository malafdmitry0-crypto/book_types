package main

import (
	"errors"
	"fmt"
)

type E struct{ N int }

func (e *E) Error() string { return fmt.Sprint(e.N) }
func main() {
	root := &E{7}
	err := fmt.Errorf("layer: %w", root)
	var found *E
	fmt.Println(errors.As(err, &found), found == root)
}
