package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x any
	json.Unmarshal([]byte(`{"n":7}`), &x)
	fmt.Printf("%T", x.(map[string]any)["n"])
}
