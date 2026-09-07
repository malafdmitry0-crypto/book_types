package main

import (
	"fmt"
)

type counter struct{ N int }

func (c *counter) Get() int { return c.N }
func (c *counter) Reset()   { c.N = 0 }

type Facade struct{ Inner *counter }

func (f Facade) Get() int { return f.Inner.Get() }
func main() {
	x := Facade{&counter{7}}
	var i interface{ Get() int } = x
	_, ok := i.(interface{ Reset() })
	fmt.Println(i.Get(), ok)
}
