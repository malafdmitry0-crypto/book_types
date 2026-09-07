package main

type Box[T any] struct{ V T }

func main() {
	var b Box
	_ = b
}
