package main

type E struct{}

func (*E) Error() string { return "E" }
func run(func() error)   {}
func main() {
	run(func() *E { return nil })
}
