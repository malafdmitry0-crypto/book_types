package main

import (
	"fmt"
	"reflect"
)

type N int

func main() {
	t := reflect.TypeFor[[]N]()
	fmt.Printf("%q %s\n", t.Name(), t.Elem().Name())
}
