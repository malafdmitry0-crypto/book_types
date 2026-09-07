package main

import (
	"fmt"
)

type E struct{}

func (e *E) Error() string {
	if e == nil {
		return "empty"
	}
	return "full"
}
func main() {
	var e error = (*E)(nil)
	fmt.Println(e.Error())
}
