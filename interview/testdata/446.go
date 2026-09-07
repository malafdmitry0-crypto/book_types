package main

import (
	"fmt"
)

type E struct{}

func (*E) Error() string { return "E" }
func each(s []int, f func(int) error) error {
	for _, v := range s {
		if e := f(v); e != nil {
			return e
		}
	}
	return nil
}
func main() {
	calls := 0
	err := each([]int{1, 2}, func(int) error {
		calls++
		var e *E
		return e
	})
	fmt.Println(err == nil, calls)
}
