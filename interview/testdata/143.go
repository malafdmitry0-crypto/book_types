package main

import (
	"fmt"
)

type T int

func (*T) Inc() {}

type O struct{ *T }

func main() {
	var i interface{ Inc() } = O{new(T)}
	i.Inc()
	fmt.Println("ok")
}
