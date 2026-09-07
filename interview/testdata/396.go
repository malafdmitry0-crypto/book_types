package main

func f[T ~int | ~int32](n T) {
	for range n {
	}
}
func main() {
	f(int32(2))
}
