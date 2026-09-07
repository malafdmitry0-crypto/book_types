package main

type Handler interface{ Handle(int) int }

func use(Handler) {}
func main() {
	use(func(n int) int { return n })
}
