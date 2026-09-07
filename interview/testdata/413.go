package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type WF func([]byte) (int, error)

func (f WF) Write(p []byte) (int, error) { return f(p) }
func main() {
	w := WF(func(p []byte) (int, error) { return len(p) - 1, nil })
	n, err := io.Copy(w, strings.NewReader("abc"))
	fmt.Println(n, errors.Is(err, io.ErrShortWrite))
}
