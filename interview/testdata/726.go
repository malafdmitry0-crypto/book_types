package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var b bytes.Buffer
	r := io.TeeReader(strings.NewReader("abcd"), &b)
	p := make([]byte, 2)
	io.ReadFull(r, p)
	fmt.Println(string(p), b.String())
}
