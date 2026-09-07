package main

import (
	"fmt"
	"reflect"
)

type N int

func Map[A, B any](s []A, f func(A) B) (out []B) {
	for _, v := range s {
		out = append(out, f(v))
	}
	return
}
func main() {
	out := Map([]N{2, 3}, func(n N) int { return int(reflect.ValueOf(n).Int()) })
	fmt.Printf("%T %v\n", out, out)
}
