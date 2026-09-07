package main

import (
	"fmt"
	"slices"
)

type MinView []int

func (s MinView) Pick() int { return slices.Min(s) }

type MaxView []int

func (s MaxView) Pick() int { return slices.Max(s) }
func main() {
	s := []int{4, 2, 8}
	fmt.Println(MinView(s).Pick(), MaxView(s).Pick())
}
