package main

import (
	"context"
	"fmt"
)

type K string

func main() {
	ctx := context.WithValue(context.Background(), K("id"), 7)
	fmt.Println(ctx.Value("id") == nil, ctx.Value(K("id")))
}
