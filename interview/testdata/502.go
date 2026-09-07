package main

type Mapper[A, B any] interface{ Map(A) B }

func main() {
	var a Mapper[int, string]
	var b Mapper[int, any] = a
	_ = b
}
