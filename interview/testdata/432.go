package main

import (
	"fmt"
)

type Fold struct {
	Values []int
	Total  int
}

func (f *Fold) Len() int  { return len(f.Values) }
func (f *Fold) Add(i int) { f.Total += f.Values[i] }
func run(f interface {
	Len() int
	Add(int)
}) {
	for i := 0; i < f.Len(); i++ {
		f.Add(i)
	}
}
func main() {
	f := Fold{Values: []int{1, 2}}
	run(&f)
	fmt.Println(f.Total)
}
