package main

import (
	"fmt"
)

type T struct{ S []int }

func (t T) Change() { t.S[0] = 9 }
func main() {
	t := T{[]int{1}}
	t.Change()
	fmt.Println(t.S)
}
