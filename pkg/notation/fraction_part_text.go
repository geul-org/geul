//ff:func feature=notation type=decoder control=selection
//ff:what fractionPartText — one FRACTION component as written: an INTEGER as digits, a DECIMAL with its point or exponent; false for an exponent-0 DECIMAL
package notation

import "github.com/geul-org/gogeul/pkg/core"

// fractionPartText writes a FRACTION numerator or denominator so that
// parseFractionPart reads it back to the same record: an INTEGER as bare
// digits, a DECIMAL by decimalText (always with a point or an exponent). A
// DECIMAL with exponent 0 (invalid, schema V10) or a broken component is
// false.
func fractionPartText(r core.Record) (string, bool) {
	switch r.Codec {
	case core.CodecSignedInteger, core.CodecUnsignedInteger:
		n, err := core.DecodeIntegerBig(r)
		if err != nil {
			return "", false
		}
		return n.String(), true
	default:
		e, c, err := core.DecodeDecimalBig(r)
		if err != nil || e == 0 {
			return "", false
		}
		return decimalText(e, c), true
	}
}
