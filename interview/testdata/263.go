package main

type Box[T any] struct{ V T }

func main() {
	var a Box[int]
	var b Box[any] = a
	_ = b
}
