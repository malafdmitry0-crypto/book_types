package main

import (
	"fmt"
)

type T int

func (t *T) Inc() { *t++ }
func main() {
	s := []T{0}
	s[0].Inc()
	fmt.Println(s[0])
}
