package main

func main() {
	var r <-chan int
	close(r)
}
