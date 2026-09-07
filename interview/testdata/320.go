package main

import (
	"fmt"
)

func count[T any](v ...T) int { return len(v) }
func main() {
	fmt.Println(count())
}
