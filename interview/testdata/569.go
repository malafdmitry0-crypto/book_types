package main

import (
	"errors"
	"fmt"
)

func main() {
	root := fmt.Errorf("root")
	wrapped := fmt.Errorf("adapter: %v", root)
	fmt.Println(errors.Is(wrapped, root))
}
