//ff:func feature=notation type=decoder control=sequence
//ff:what formatDecimal — a SIGNED-DECIMAL as a bare number, an UNSIGNED-DECIMAL by name, digits as written
package notation

import (
	"math/big"

	"github.com/geul-org/gogeul/pkg/core"
)

// formatDecimal writes a decimal Value (decimalText). SIGNED-DECIMAL is the
// bare number; UNSIGNED-DECIMAL (coefficient up to int64) is
// UNSIGNED-DECIMAL(n).
func formatDecimal(r core.Record) (string, bool) {
	if r.Codec == core.CodecSignedDecimal {
		e, c, err := core.DecodeSignedDecimalBig(r)
		if err != nil {
			return "", false
		}
		return decimalText(e, c), true
	}
	e, c, err := core.DecodeDecimalValue(r)
	if err != nil {
		return "", false
	}
	return "UNSIGNED-DECIMAL(" + decimalText(e, big.NewInt(c)) + ")", true
}
