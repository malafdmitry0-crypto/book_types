package main

import (
	"fmt"
	"io"
	"strings"
)

type OnlyRead struct{ io.Reader }

func main() {
	src := io.NopCloser(strings.NewReader("go"))
	var r io.Reader = OnlyRead{src}
	_, ok := r.(io.Closer)
	fmt.Println(ok)
}
