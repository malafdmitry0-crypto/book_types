package main

import (
	"fmt"
)

func wrap(name string, next func(), log *[]string) func() {
	return func() {
		*log = append(*log, name+"+")
		next()
		*log = append(*log, name+"-")
	}
}
func main() {
	var log []string
	f := wrap("A", wrap("B", func() { log = append(log, "X") }, &log), &log)
	f()
	fmt.Println(log)
}
