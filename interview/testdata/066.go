package main

import (
	"fmt"
)

type T int

func main() {
	a := struct{ T }{1}
	var b struct{ T T } = a
	fmt.Println(b)
}
