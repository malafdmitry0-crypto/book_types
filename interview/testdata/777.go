package main

import (
	"fmt"
)

func main() {
	x, y := 1, 2
	s := []*int{&x, &y}
	a := [2]*int(s)
	s[0] = &y
	*a[0] = 9
	fmt.Println(x, *s[0], *a[0], a[0] == s[0])
}
