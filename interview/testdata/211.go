package main

import (
	"fmt"
)

func main() {
	m := map[any]string{int(1): "int", int64(1): "int64"}
	fmt.Println(len(m), m[1], m[int64(1)])
}
