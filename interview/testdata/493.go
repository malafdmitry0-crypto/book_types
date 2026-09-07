package main

func Group[E any, K comparable](s []E, key func(E) K) map[K][]E { return nil }
func main() {
	_ = Group([]int{1}, func(n int) []int { return []int{n} })
}
