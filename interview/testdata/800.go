package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Compose[A, B, C any](f func(A) B, g func(B) C) func(A) C { return func(x A) C { return g(f(x)) } }
func main() {
	f := reflect.ValueOf(Compose[int, string, bool](func(x int) string { return strings.Repeat("x", x) }, func(s string) bool { return len(s) > 2 }))
	r := f.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(r[0].Bool(), f.Type().NumIn())
}
