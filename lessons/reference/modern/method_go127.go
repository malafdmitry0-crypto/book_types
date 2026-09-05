//go:build go1.27

package modern
type Converter struct{}

func (Converter) ToInt64[T ~int | ~int32 | ~int64](x T) int64 {
    return int64(x)
}

// n := (Converter{}).ToInt64(int32(42)) // Go 1.27+
