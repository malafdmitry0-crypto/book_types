package main

import (
	"fmt"
)

type T struct{}

func (T) F() string { return "inner" }

type O struct{ T }

func (O) F() string { return "outer" }
func main() {
	fmt.Println(O{}.F(), O{}.T.F())
}
