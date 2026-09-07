package main

type T int

func (t *T) Inc() {}
func makeT() T    { return 0 }
func main() {
	makeT().Inc()
}
