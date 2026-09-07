package main

import (
	"fmt"
)

type View struct{ S []int }

func (v View) Len() int           { return len(v.S) }
func (v View) Less(i, j int) bool { return v.S[i] < v.S[j] }
func pick(v interface {
	Len() int
	Less(int, int) bool
}) int {
	best := 0
	for i := 1; i < v.Len(); i++ {
		if v.Less(i, best) {
			best = i
		}
	}
	return best
}
func main() {
	s := []int{2, 1, 1}
	fmt.Println(pick(View{s}))
}
