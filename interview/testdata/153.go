package main

type I interface{ F() }
type O struct{ X I }

func main() {
	var o O
	var i interface{ F() } = o
	_ = i
}
