package main

import (
	"fmt"
)

type T struct{ S []int }

func (t T) Clear() { t.S = nil }
func main() {
	t := T{[]int{1}}
	t.Clear()
	fmt.Println(len(t.S))
}
