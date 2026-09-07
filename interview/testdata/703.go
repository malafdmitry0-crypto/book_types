package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := 7
	err := json.Unmarshal([]byte(`null`), &x)
	fmt.Println(x, err == nil)
}
