package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var x struct{ N int }
	json.Unmarshal([]byte(`{"N":1,"N":8}`), &x)
	fmt.Println(x.N)
}
