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
	var p *E
	var e error = p
	fmt.Println(e == nil, e.Error())
}
