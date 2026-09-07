package main

type A struct{ X int }
type O struct{ A }

func main() {
	_ = O{X: 1}
}
