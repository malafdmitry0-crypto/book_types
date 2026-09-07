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
	out, err := MapErr([]int{1, -1, 3}, func(n int) (string, error) {
		if n < 0 {
			return "bad", fmt.Errorf("negative")
		}
		return fmt.Sprint(n), nil
	})
	fmt.Println(out, err != nil)
}
