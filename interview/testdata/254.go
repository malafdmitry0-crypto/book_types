package main

func main() {
	var c chan int
	var a chan any = c
	_ = a
}
