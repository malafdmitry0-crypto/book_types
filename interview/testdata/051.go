package main

type P *int

func (P) F() {}
func main() {
	var p P
	_ = p
}
