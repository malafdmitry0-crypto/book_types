package main

import (
	"fmt"
)

func mapInts(s []int, f func(int) int) (out []int) {
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}
func main() {
	var log []string
	f := func(n int) int {
		log = append(log, "f")
		return n + 1
	}
	g := func(n int) int {
		log = append(log, "g")
		return n * 2
	}
	_ = mapInts(mapInts([]int{1, 2}, f), g)
	fmt.Println(log)
	log = nil
	_ = mapInts([]int{1, 2}, func(n int) int { return g(f(n)) })
	fmt.Println(log)
}
