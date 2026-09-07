package main

import (
	"errors"
	"fmt"
)

func Adapt(f func() (int, error)) func() (int, error) {
	return func() (int, error) {
		v, e := f()
		if e != nil {
			return v, fmt.Errorf("adapt: %w", e)
		}
		return v, nil
	}
}
func main() {
	f := Adapt(func() (int, error) { return 7, errors.New("x") })
	v, e := f()
	fmt.Println(v, e != nil)
}
