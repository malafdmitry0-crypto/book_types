package main

type Item int

func (i Item) Less(j Item) bool { return i < j }
func main() {
	var x interface{ Less(int, int) bool } = Item(3)
	_ = x
}
