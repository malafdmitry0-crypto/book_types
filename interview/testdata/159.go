package main

type I interface{ F() }
type O struct{ *I }

func main() {
	var o O
	_ = o
}
