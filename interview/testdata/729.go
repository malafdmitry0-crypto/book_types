package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := io.NopCloser(strings.NewReader("go"))
	r.Close()
	b, e := io.ReadAll(r)
	fmt.Println(string(b), e == nil)
}
