package main

import (
	"fmt"
)

type T int

func (t T) Inc() { t++ }
func main() {
	t := T(1)
	t.Inc()
	fmt.Println(t)
}
