package main

import (
	"fmt"
)

type ID int

func First[S []E, E any](s S) E { return s[0] }
func main() {
	fmt.Println(First([]ID{7}))
}
