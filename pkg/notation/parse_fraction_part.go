//ff:func feature=notation type=encoder control=sequence
//ff:what parseFractionPart — one FRACTION component: digits as an INTEGER, a number with a point or exponent as a DECIMAL (exponent 0 refused); signed for the numerator
package notation

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseFractionPart reads a FRACTION component as written
// (value/numeric/FRACTION 표기 보존): plain digits are an INTEGER record, a
// number with a point or an exponent a DECIMAL record keeping its digits —
// SIGNED-* for the numerator (signed), UNSIGNED-* for the denominator. A
// DECIMAL whose exponent comes out 0 ("5.", "5e0") is refused: an integer
// original is an INTEGER.
func parseFractionPart(arg string, signed bool) (core.Record, error) {
	if !strings.ContainsAny(arg, ".eE") {
		n, ok := new(big.Int).SetString(arg, 10)
		if !ok {
			return core.Record{}, fmt.Errorf("geul/notation: %q is not an integer", arg)
		}
		return core.EncodeIntegerBig(signed, n)
	}
	e, c, err := parseDecimal(arg)
	if err != nil {
		return core.Record{}, err
	}
	if e == 0 {
		return core.Record{}, fmt.Errorf("geul/notation: %q has exponent 0; write the integer", arg)
	}
	return core.EncodeDecimalBig(signed, e, c)
}
