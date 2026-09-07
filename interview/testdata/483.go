package main

import (
	"fmt"
)

type Names []string

func Map[Out ~[]R, E, R any](s []E, f func(E) R) Out { return nil }
func main() {
	var out Names = Map([]int{1}, func(n int) string { return fmt.Sprint(n) })
	fmt.Println(out)
}
