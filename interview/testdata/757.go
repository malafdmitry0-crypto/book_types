package main

import (
	"context"
	"fmt"
)

type key struct{}

func main() {
	base := context.WithValue(context.Background(), key{}, 7)
	ctx, cancel := context.WithCancel(base)
	safe := context.WithoutCancel(ctx)
	cancel()
	fmt.Println(safe.Done() == nil, safe.Err() == nil, safe.Value(key{}))
}
