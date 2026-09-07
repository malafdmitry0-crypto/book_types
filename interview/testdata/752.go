package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	x := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		x = 5
	}()
	wg.Wait()
	fmt.Println(x)
}
