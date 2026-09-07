package main

type T int

func (*T) Inc() {}

type O struct{ T }

func main() {
	var i interface{ Inc() } = O{}
	_ = i
}
