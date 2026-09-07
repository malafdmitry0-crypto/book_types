package main

import (
	"fmt"
)

type B byte

func main() {
	b := []B{65, 66}
	fmt.Println(string(b))
}
