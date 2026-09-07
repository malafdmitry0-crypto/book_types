package main

import (
	"errors"
	"fmt"
)

func main() {
	root := fmt.Errorf("root")
	wrapped := fmt.Errorf("adapter: %w", root)
	fmt.Println(errors.Is(wrapped, root), wrapped == root)
}
