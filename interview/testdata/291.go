package main

type C interface{ ~int | ~string }

func main() {
	var x C
	_ = x
}
