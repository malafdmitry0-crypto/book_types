package main

import (
	"fmt"
)

func around(next func(), log *[]string) func() {
	return func() {
		*log = append(*log, "enter")
		defer func() { *log = append(*log, "leave") }()
		next()
	}
}
func main() {
	var log []string
	f := around(func() { log = append(log, "body") }, &log)
	f()
	fmt.Println(log)
}
