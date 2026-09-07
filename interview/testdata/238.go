package main

import (
	"fmt"
)

func pair() (int, int) { return 2, 3 }
func main() {
	fmt.Println(1, pair())
}
