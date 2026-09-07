package main

import (
	"fmt"
)

func main() {
	s := "я"
	fmt.Printf("%T %d %d\n", s[0], s[0], []rune(s)[0])
}
