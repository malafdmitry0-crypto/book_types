package main

import (
	"fmt"
)

type View struct{ Values *[]int }

func (v View) Len() int { return len(*v.Values) }
func main() {
	s := []int{1}
	a := View{&s}
	s = append(s, 2)
	fmt.Println(a.Len())
}
