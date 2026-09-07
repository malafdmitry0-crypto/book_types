package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Trace struct {
	Inner io.Reader
	Calls *int
}

func (t Trace) Read(p []byte) (int, error) {
	*t.Calls++
	return t.Inner.Read(p)
}
func main() {
	calls := 0
	src := Trace{Inner: strings.NewReader("go"), Calls: &calls}
	var dst bytes.Buffer
	io.Copy(&dst, src)
	fmt.Println(dst.String(), calls > 0)
}
