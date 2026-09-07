package main

type T int

func (t *T) Inc() {}
func main() {
	m := map[int]T{0: 0}
	m[0].Inc()
}
