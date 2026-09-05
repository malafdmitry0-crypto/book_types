package invalid
type Stringer interface {String()string}
func f(v interface{}) {var _ Stringer=v}
