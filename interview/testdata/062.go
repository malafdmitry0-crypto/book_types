package main

import (
	"fmt"
)

func main() {
	a := struct{ X int }{1}
	var b struct{ Y int } = a
	fmt.Println(b)
}
