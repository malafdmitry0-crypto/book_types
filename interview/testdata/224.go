package main

func f(...any) {}
func main() {
	a := []int{1, 2}
	f(a...)
}
