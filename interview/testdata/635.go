package main

import (
	"fmt"
)

func main() {
	goto Done
	x := 2
Done:
	fmt.Println(x)
}
