package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	a := errors.New("a")
	cancel(a)
	cancel(errors.New("b"))
	fmt.Println(context.Cause(ctx) == a, ctx.Err() == context.Canceled)
}
