package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	d := json.NewDecoder(strings.NewReader(`{"N":4,"extra":8}`))
	d.DisallowUnknownFields()
	var x struct{ N int }
	err := d.Decode(&x)
	fmt.Println(err != nil, x.N)
}
