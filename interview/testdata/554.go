package main

import (
	"fmt"
	"reflect"
)

func main() {
	b := make([]byte, 3)
	n := reflect.Copy(reflect.ValueOf(b), reflect.ValueOf("яx"))
	fmt.Println(n, b)
}
