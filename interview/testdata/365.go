package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "я"
	fmt.Println(len(s[:1]), utf8.ValidString(s[:1]))
}
