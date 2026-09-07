package main

type F func() int

func (f F) Read() int { return f() }
func main() {
	var r interface{ Read() any } = F(func() int { return 1 })
	_ = r
}
