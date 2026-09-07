package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b, _ := json.Marshal([]byte{71, 111})
	fmt.Println(string(b))
}
