package main

import (
	"fmt"
)

func fold(s []int, f func(int, int) (int, error)) (a int, e error) {
	for _, v := range s {
		a, e = f(a, v)
		if e != nil {
			return
		}
	}
	return
}
func main() {
	n, err := fold([]int{1, 2}, func(a, v int) (int, error) {
		if v == 2 {
			return 99, fmt.Errorf("stop")
		}
		return a + v, nil
	})
	fmt.Println(n, err != nil)
}
