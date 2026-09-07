package main

import (
	"fmt"
)

func fold(s []int, f func(int, int) (int, error)) (int, error) {
	a := 0
	for _, v := range s {
		n, e := f(a, v)
		if e != nil {
			return a, e
		}
		a = n
	}
	return a, nil
}
func main() {
	events := []int{}
	sum, err := fold([]int{1, 2, 3}, func(a, v int) (int, error) {
		events = append(events, v)
		if v == 2 {
			return 99, fmt.Errorf("stop")
		}
		return a + v, nil
	})
	fmt.Println(sum, err != nil, events)
}
