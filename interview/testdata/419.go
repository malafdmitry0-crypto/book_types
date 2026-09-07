package main

type F func(...int) int

func (f F) Apply(x ...int) int { return f(x...) }
func main() {
	var r interface{ Apply([]int) int } = F(func(x ...int) int { return len(x) })
	_ = r
}
