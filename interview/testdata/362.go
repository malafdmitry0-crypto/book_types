package main

import (
	"fmt"
)

func main() {
	r := rune(0xD800)
	fmt.Printf("%q %d\n", string(r), len(string(r)))
}
