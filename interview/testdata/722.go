package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	n, e := io.ReadFull(strings.NewReader(""), make([]byte, 2))
	fmt.Println(n, e == io.EOF, e == io.ErrUnexpectedEOF)
}
