package main

func main() {
	var s chan<- int
	_ = <-s
}
