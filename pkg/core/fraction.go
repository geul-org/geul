//ff:type feature=core type=model
//ff:what Fraction — the two records of a FRACTION Value: numerator (SIGNED-INTEGER or SIGNED-DECIMAL) and denominator (UNSIGNED-INTEGER or UNSIGNED-DECIMAL)
package core

// Fraction is the body of a FRACTION Value (value/numeric/FRACTION): the
// numerator and the denominator as whole Value records, kept as written (not
// reduced). An integer component of the original is an INTEGER record and a
// component with a decimal point a DECIMAL record; the sign is on the
// numerator only.
type Fraction struct {
	Numerator   Record
	Denominator Record
}
