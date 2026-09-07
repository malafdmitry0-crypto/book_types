package main

type Logger interface{ Log(...any) }
type Wrap struct{}

func (Wrap) Log([]any) {}
func main() {
	var i Logger = Wrap{}
	_ = i
}
