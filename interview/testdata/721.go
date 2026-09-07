package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	b := make([]byte, 4)
	n, e := io.ReadFull(strings.NewReader("go"), b)
	fmt.Println(n, e == io.ErrUnexpectedEOF, string(b[:n]))
}
