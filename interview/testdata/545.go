package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := reflect.ValueOf([]any{int32(7)})
	fmt.Println(s.Index(0).Kind(), s.Index(0).Elem().Kind())
}
