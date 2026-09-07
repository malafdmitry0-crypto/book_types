package main

func main() {
	var f func(...int)
	var g func([]int) = f
	_ = g
}
