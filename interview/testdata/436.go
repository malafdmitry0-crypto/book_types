package main

import (
	"fmt"
)

type View struct{ Values []int }

func (v View) Len() int { return len(v.Values) }
func main() {
	s := make([]int, 1, 3)
	a := View{s}
	s = append(s, 9)
	fmt.Println(a.Len(), a.Values[:2])
}
