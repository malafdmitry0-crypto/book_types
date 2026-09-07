package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var n atomic.Int64
	n.Store(4)
	old := n.Swap(9)
	ok := n.CompareAndSwap(4, 12)
	fmt.Println(old, ok, n.Load())
}
