package main

import (
	"fmt"
)

func main() {
	r := append([]rune{}, "я"...)
	fmt.Println(r)
}
