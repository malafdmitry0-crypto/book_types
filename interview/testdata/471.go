package main

import (
	"fmt"
)

func then(f, g func(int) (int, error)) func(int) (int, error) {
	return func(n int) (int, error) {
		v, e := f(n)
		if e != nil {
			return v, e
		}
		return g(v)
	}
}
func main() {
	calls := 0
	f := func(int) (int, error) { return 0, fmt.Errorf("stop") }
	g := func(n int) (int, error) {
		calls++
		return n, nil
	}
	_, e := then(f, g)(1)
	fmt.Println(e != nil, calls)
}
