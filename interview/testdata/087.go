package main

import (
	"fmt"
)

type Alias = int

func main() {
	var a any = int(7)
	n, ok := a.(Alias)
	fmt.Println(n, ok)
}
