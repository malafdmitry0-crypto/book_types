package main

func New[E any, M interface{ Map(E) E }]() M {
	var m M
	return m
}
func main() {
	_ = New[int]()
}
