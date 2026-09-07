package main

import (
	"fmt"
)

type I interface{ Get() int }
type V int

func (v V) Get() int { return int(v) }

type Wrapper struct{ I }

func main() {
	w := Wrapper{I: V(1)}
	f := w.Get
	w.I = V(2)
	fmt.Println(f(), w.Get())
}
