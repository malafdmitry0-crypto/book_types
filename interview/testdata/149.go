package main

import (
	"fmt"
)

type T struct{}

func (T) F() string   { return "inner" }
func (t T) G() string { return t.F() }

type O struct{ T }

func (O) F() string { return "outer" }
func main() {
	fmt.Println(O{}.G())
}
