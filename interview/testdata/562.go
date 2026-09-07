package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Trace struct {
	*strings.Reader
	Calls *int
}

func (t Trace) Read(p []byte) (int, error) {
	*t.Calls++
	return t.Reader.Read(p)
}
func main() {
	calls := 0
	var src io.Reader = Trace{strings.NewReader("go"), &calls}
	var dst bytes.Buffer
	io.Copy(&dst, src)
	fmt.Println(calls)
}
