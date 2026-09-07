package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	n := 0
	once.Do(func() { n += 2 })
	once.Do(func() { n += 10 })
	fmt.Println(n)
}
