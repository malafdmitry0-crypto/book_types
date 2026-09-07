package main

import (
	"fmt"
	"io"
	"strings"
)

type Wrapped struct{ io.ReadCloser }

func main() {
	src := io.NopCloser(strings.NewReader("go"))
	var r io.Reader = Wrapped{src}
	_, ok := r.(io.Closer)
	fmt.Println(ok)
}
