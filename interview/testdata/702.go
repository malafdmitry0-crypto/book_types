package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	d := json.NewDecoder(strings.NewReader(`1234567890123456789`))
	d.UseNumber()
	var x any
	d.Decode(&x)
	fmt.Printf("%T %v", x, x)
}
