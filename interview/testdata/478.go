package main

import (
	"fmt"
)

func choose(ok bool) func() {
	if ok {
		return func() {}
	}
	return nil
}
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	f := choose(false)
	f()
}
