package main

type A struct{}

func (A) F() {}

type B struct{}

func (B) F() {}

type O struct {
	A
	B
}

func main() {
	O{}.F()
}
