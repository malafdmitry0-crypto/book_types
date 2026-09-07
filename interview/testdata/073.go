package main

type A int
type B int

func main() {
	var a *A
	var b *B = a
	_ = b
}
