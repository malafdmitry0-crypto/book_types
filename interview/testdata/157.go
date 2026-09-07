package main

type A interface{ F() int }
type B interface{ F() string }
type C interface {
	A
	B
}

func main() {
	var c C
	_ = c
}
