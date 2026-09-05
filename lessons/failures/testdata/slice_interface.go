package invalid
import("gotypes/fraction")
type Stringer interface {String()string}
var _ []Stringer=[]fraction.Fraction{}
