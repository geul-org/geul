//ff:func feature=core type=decoder control=selection
//ff:what NumberRat — the exact value of an INTEGER, DECIMAL or FRACTION Value as a rational number (the comparison layer)
package core

import (
	"fmt"
	"math/big"
)

// NumberRat returns the exact value of r for comparison
// (value/numeric/FRACTION 비교): a SIGNED-/UNSIGNED-INTEGER, a
// SIGNED-/UNSIGNED-DECIMAL C×10^E, or a FRACTION (FractionRat). Byte-different
// records of one value (2/4, 1/2, 0.5) give equal results. Other codecs fail.
func NumberRat(r Record) (*big.Rat, error) {
	switch {
	case r.Kind != KindValue:
		return nil, fmt.Errorf("geul/core: kind %d is not a numeric Value", r.Kind)
	case r.Codec == CodecSignedInteger || r.Codec == CodecUnsignedInteger:
		n, err := DecodeIntegerBig(r)
		if err != nil {
			return nil, err
		}
		return new(big.Rat).SetInt(n), nil
	case r.Codec == CodecSignedDecimal || r.Codec == CodecUnsignedDecimal:
		e, c, err := DecodeDecimalBig(r)
		if err != nil {
			return nil, err
		}
		return decimalRat(e, c), nil
	case r.Codec == CodecFraction:
		return FractionRat(r)
	default:
		return nil, fmt.Errorf("geul/core: codec 0x%02X has no exact rational value", r.Codec)
	}
}
