package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := reflect.ValueOf(func(n int) (string, bool) { return fmt.Sprint(n), n > 0 })
	out := f.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(len(out), out[0].String(), out[1].Bool())
}
