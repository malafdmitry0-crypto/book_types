package main

import (
	"fmt"
)

func protect(next func()) func() (failed bool) {
	return func() (failed bool) {
		defer func() { failed = recover() != nil }()
		next()
		return false
	}
}
func main() {
	f := protect(func() { panic("bad") })
	fmt.Println(f())
}
