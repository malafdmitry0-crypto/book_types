package main

type Op[E any] interface{ Apply(E, E) E }
type Plus struct{}

func (Plus) Apply(a, b int) int { return a + b }
func main() {
	s := []Plus{{}}
	var ops []Op[int] = s
	_ = ops
}
