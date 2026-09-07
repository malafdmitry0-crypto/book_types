package main

import (
	"fmt"
)

type S struct{ X int }

func main() {
	s := S{1}
	var a any = s
	s.X = 2
	fmt.Println(a.(S).X, s.X)
}
