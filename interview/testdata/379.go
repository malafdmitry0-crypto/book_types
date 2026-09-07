package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "é"
	fmt.Println(len(s), utf8.RuneCountInString(s), s == "é")
}
