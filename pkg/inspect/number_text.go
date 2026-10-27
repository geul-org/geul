//ff:func feature=inspect type=decoder control=selection
//ff:what numberText — a numeric Value as exact text: an integer's digits, a decimal with its quantum, a BINARY-FLOAT in shortest form, a FRACTION as n/d
package inspect

import (
	"fmt"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// numberText writes the numeric Values a QUANTITY or FRACTION holds without
// rounding: SIGNED-/UNSIGNED-INTEGER as digits (any size), SIGNED-/UNSIGNED-
// DECIMAL by decimalString (1.20 stays 1.20), BINARY-FLOAT up to binary64 in
// its shortest round-trip form, and FRACTION as numerator/denominator, each
// component written as it is. Other codecs and broken payloads fail.
func numberText(r core.Record) (string, error) {
	switch r.Codec {
	case core.CodecSignedInteger, core.CodecUnsignedInteger:
		n, err := core.DecodeIntegerBig(r)
		if err != nil {
			return "", err
		}
		return n.String(), nil
	case core.CodecSignedDecimal, core.CodecUnsignedDecimal:
		e, c, err := core.DecodeDecimalBig(r)
		if err != nil {
			return "", err
		}
		return decimalString(e, c), nil
	case core.CodecBinaryFloat:
		f, err := binaryFloat(r.Payload)
		return strconv.FormatFloat(f, 'g', -1, 64), err
	case core.CodecFraction:
		return fractionText(r)
	default:
		return "", fmt.Errorf("geul/inspect: codec 0x%02X is not a number", r.Codec)
	}
}
