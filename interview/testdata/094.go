package main

type R interface{ Read() }
type T struct{}

func main() {
	var r R
	_ = r.(T)
}
