package main

type T int

func (*T) Inc() {}
func main() {
	_ = T.Inc
}
