package main

import (
	"fmt"
)

func main() {
	x := 1
	s := []*int{&x}
	a := [1]*int(s)
	*a[0] = 9
	fmt.Println(*s[0])
}
