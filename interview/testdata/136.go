package main

import (
	"fmt"
)

type T int

func (t T) Get() int { return 1 }
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var t *T
	fmt.Println(t.Get())
}
