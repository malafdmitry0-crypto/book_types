package main

import (
	"fmt"
)

func main() {
	var n int64 = 9007199254740993
	fmt.Println(int64(float64(n)) == n)
}
