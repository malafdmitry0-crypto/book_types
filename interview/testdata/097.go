package main

type S struct{ X int }

func main() {
	var a any = S{1}
	a.(S).X = 2
}
