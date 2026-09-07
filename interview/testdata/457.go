package main

import (
	"fmt"
)

type Item struct{ N int }

func mapPtrs(s []int, f func(int) *Item) (out []*Item) {
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}
func main() {
	box := Item{}
	out := mapPtrs([]int{1, 2}, func(n int) *Item {
		box.N = n
		return &box
	})
	fmt.Println(out[0] == out[1], out[0].N)
}
