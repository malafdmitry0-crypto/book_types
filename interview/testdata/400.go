package main

func main() {
	s := func(yield func(int) int) { yield(1) }
	for range s {
	}
}
