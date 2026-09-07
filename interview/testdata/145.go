package main

import (
	"fmt"
)

type T int

func (T) Get() int { return 1 }

type O struct{ *T }

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var o O
	fmt.Println(o.Get())
}
