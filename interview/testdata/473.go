package main

import (
	"fmt"
)

func factory() func() {
	defer func() { _ = recover() }()
	return func() { panic("later") }
}
func main() {
	defer func() { fmt.Println("outer", recover() != nil) }()
	f := factory()
	f()
}
