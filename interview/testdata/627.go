package main

import (
	"fmt"
)

func main() {
	var ps []*int
	for _, v := range []int{3, 8} {
		ps = append(ps, &v)
	}
	fmt.Println(ps[0] == ps[1], *ps[0], *ps[1])
}
