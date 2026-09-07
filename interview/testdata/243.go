package main

func main() {
	var r <-chan int
	_ = (chan int)(r)
}
