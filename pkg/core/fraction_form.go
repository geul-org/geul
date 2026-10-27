//ff:func feature=core type=decoder control=selection
//ff:what fractionForm — whether the two FRACTION records have their codecs: SIGNED-INTEGER or SIGNED-DECIMAL numerator, UNSIGNED-INTEGER or UNSIGNED-DECIMAL denominator
package core

import "errors"

// fractionForm checks the codecs of a FRACTION's numerator and denominator
// records (value/numeric/FRACTION). The list is closed: no other Value, no
// nested FRACTION and no slot marker.
func fractionForm(n, d Record) error {
	switch {
	case n.Kind != KindValue || n.Codec != CodecSignedInteger && n.Codec != CodecSignedDecimal:
		return errors.New("geul/core: FRACTION numerator is not a SIGNED-INTEGER or SIGNED-DECIMAL Value")
	case d.Kind != KindValue || d.Codec != CodecUnsignedInteger && d.Codec != CodecUnsignedDecimal:
		return errors.New("geul/core: FRACTION denominator is not an UNSIGNED-INTEGER or UNSIGNED-DECIMAL Value")
	default:
		return nil
	}
}
