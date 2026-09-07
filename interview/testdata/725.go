package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := io.MultiReader(strings.NewReader("a"), strings.NewReader(""), strings.NewReader("bc"))
	b, e := io.ReadAll(r)
	fmt.Println(string(b), e == nil)
}
