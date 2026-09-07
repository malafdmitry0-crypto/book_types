package main

import (
	"fmt"
)

type S struct{ X int }

func main() {
	m := map[int]S{1: {2}}
	v := m[1]
	v.X = 9
	m[1] = v
	fmt.Println(m[1].X)
}
