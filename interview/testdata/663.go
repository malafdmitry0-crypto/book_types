package main

import (
	"fmt"
)

type IDs []int

func First[S []E, E any](s S) E { return s[0] }
func main() {
	fmt.Println(First(IDs{7}))
}
