package main

import (
	"fmt"
)

type T int

func (t *T) IsNil() bool { return t == nil }
func main() {
	var t *T
	fmt.Println(t.IsNil())
}
