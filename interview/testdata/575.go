package main

import (
	"errors"
	"fmt"
)

type Code int

func (c Code) Error() string { return fmt.Sprint(int(c)) }
func main() {
	root := Code(7)
	fmt.Println(errors.Is(root, Code(7)), errors.Is(root, Code(8)))
}
