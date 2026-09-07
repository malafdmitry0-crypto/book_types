package main

type Alias = int

func main() {
	var x any = 1
	switch x.(type) {
	case int:
	case Alias:
	}
}
