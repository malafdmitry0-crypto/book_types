package main

import (
	"fmt"
)

type T int
type Alias = T
type O struct{ Alias }

func main() {
	o := O{Alias: T(3)}
	fmt.Println(o.Alias)
}
