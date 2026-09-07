package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a, _ := json.Marshal([]int(nil))
	b, _ := json.Marshal([]int{})
	fmt.Println(string(a), string(b))
}
