package main

type Step[E any] interface{ Add(E) }
type Acc struct{ N int }

func (a *Acc) Add(n int) { a.N += n }
func main() {
	var op Step[int] = Acc{}
	_ = op
}
