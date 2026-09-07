package main

import (
	"fmt"
)

type I interface{ F() }
type O struct{ I }

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var o O
	o.F()
}
