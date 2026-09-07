package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := reflect.TypeFor[chan int]()
	b := reflect.TypeFor[<-chan int]()
	fmt.Println(a == b, a.ChanDir(), b.ChanDir())
}
