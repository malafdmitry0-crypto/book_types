package main

type T int

func (t *T) Inc() { *t++ }
func main() {
	var t T
	var i interface{ Inc() } = t
	_ = i
}
