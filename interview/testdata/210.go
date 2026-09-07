package main

import (
	"fmt"
)

type S struct{ X int }

func main() {
	m := map[int]*S{1: {2}}
	m[1].X = 9
	fmt.Println(m[1].X)
}
