package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	x := 0
	go func() {
		x = 7
		close(done)
	}()
	<-done
	fmt.Println(x)
}
