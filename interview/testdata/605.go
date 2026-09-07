package main

import (
	"fmt"
)

type N struct{ V int }

func (n N) Show() { fmt.Println(n.V) }
func main() {
	p := &N{2}
	defer p.Show()
	p.V = 8
}
