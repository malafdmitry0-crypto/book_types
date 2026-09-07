package main

import (
	"fmt"
)

type E struct{}

func (*E) Error() string { return "E" }
func main() {
	f := func() *E { return nil }
	cb := func() error {
		if e := f(); e != nil {
			return e
		}
		return nil
	}
	fmt.Println(cb() == nil)
}
