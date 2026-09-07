package main

import (
	"fmt"
)

type Factory interface{ Make() any }
type N struct{}

func (N) Make() N { return N{} }
func main() {
	var f Factory = N{}
	fmt.Println(f.Make())
}
