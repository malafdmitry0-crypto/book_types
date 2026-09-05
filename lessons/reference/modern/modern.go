package modern

func Identity[T any](x T) T { return x }

var intIdentity func(int) int = Identity
