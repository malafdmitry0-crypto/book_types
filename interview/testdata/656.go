package main

import (
	"fmt"
)

type N struct{}

func (N) M() { fmt.Println(7) }

type Wrapper struct{ *N }

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var w Wrapper
	w.M()
}
