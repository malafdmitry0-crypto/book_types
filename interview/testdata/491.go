package main

import (
	"fmt"
)

func MapErr[E, R any](s []E, f func(E) (R, error)) ([]R, error) {
	out := make([]R, 0, len(s))
	for _, v := range s {
		r, e := f(v)
		if e != nil {
			return out, e
		}
		out = append(out, r)
	}
	return out, nil
}
func main() {
	out, _ := MapErr([]int{1, -1, 3}, func(n int) (int, error) {
		if n < 0 {
			return 0, fmt.Errorf("stop")
		}
		return n, nil
	})
	fmt.Println(len(out), cap(out))
}
