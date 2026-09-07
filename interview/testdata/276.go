package main

import (
	"fmt"
)

type Box[T any] struct{ V T }

func (b Box[int]) Add() int { return b.V + 1 }
func main() {
	fmt.Println(Box[int]{3}.Add())
}
