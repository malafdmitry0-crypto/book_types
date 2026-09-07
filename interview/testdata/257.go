package main

type C chan int
type R <-chan int

func main() {
	var c C
	var r R = c
	_ = r
}
