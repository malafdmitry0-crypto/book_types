package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("abcdef")
	a, _ := io.ReadAll(io.LimitReader(r, 3))
	b, _ := io.ReadAll(r)
	fmt.Println(string(a), string(b))
}
