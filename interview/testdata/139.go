package main

type T struct{}

func (T) Clone() T { return T{} }
func main() {
	var x interface{ Clone() any } = T{}
	_ = x
}
