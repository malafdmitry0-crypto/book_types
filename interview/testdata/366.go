package main

import (
	"fmt"
)

func main() {
	s := "\xff\xff"
	fmt.Println(len(s), len([]rune(s)), len(string([]rune(s))))
}
