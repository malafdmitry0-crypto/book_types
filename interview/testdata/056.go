package main

type F func(int) int
type G func(int) int

func main() {
	var f F
	var g G = f
	_ = g
}
