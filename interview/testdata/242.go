package main

func main() {
	var r <-chan int
	var c chan int = r
	_ = c
}
