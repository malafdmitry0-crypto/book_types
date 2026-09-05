package invalid
type Stringer interface {String()string}
func f(v Stringer){_ = v.Numerator}
