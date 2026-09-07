package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	func() {
		defer func() { recover() }()
		once.Do(func() { panic("x") })
	}()
	n := 0
	once.Do(func() { n++ })
	fmt.Println(n)
}
