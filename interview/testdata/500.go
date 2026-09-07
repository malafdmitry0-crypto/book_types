package main

import (
	"fmt"
)

func main() {
	_ = []func(any) any{func(n int) string { return fmt.Sprint(n) }}
}
