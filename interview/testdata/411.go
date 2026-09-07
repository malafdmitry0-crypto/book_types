package main

import (
	"fmt"
)

type F func()

func (f F) Run() { f() }
func decorate(next interface{ Run() }, log *[]string) F {
	return func() {
		*log = append(*log, "before")
		next.Run()
		*log = append(*log, "after")
	}
}
func main() {
	var log []string
	h := decorate(F(func() { log = append(log, "body") }), &log)
	h.Run()
	fmt.Println(log)
}
